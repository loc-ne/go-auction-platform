package usecase

import (
	"context"
    "time"
    "fmt"
    "log"
	"github.com/loc-ne/go-auction/services/product/internal/entity"
	repoRedis "github.com/loc-ne/go-auction/services/product/internal/repository/redis"
    "github.com/redis/go-redis/v9"
)

type ProductRepository interface {
    Create(ctx context.Context, product *entity.Product) error
    GetActiveAuctions(ctx context.Context, limit, offset int) ([]entity.Product, error)
    GetByID(ctx context.Context, id string) (*entity.Product, error)
}

type ProductUsecase interface {
	Create(ctx context.Context, product *entity.Product) error
	GetActiveAuctions(ctx context.Context, limit, offset int) ([]entity.Product, error)
	GetByID(ctx context.Context, id string) (*entity.Product, error)
}

type productUsecase struct {
    repo ProductRepository
	redisClient *repoRedis.RedisClient 
}

func NewProductUsecase(r ProductRepository, red *repoRedis.RedisClient) ProductUsecase {
    return &productUsecase{
        repo:      r,
		redisClient: red,
    }
}

func (u *productUsecase) Create(ctx context.Context, product *entity.Product) error {
    err := u.repo.Create(ctx, product)
    if err != nil {
        return err
    }

    if product.Status == 1 {
        ttl := time.Until(product.EndAt)
        
        if ttl > 0 {
            priceKey := fmt.Sprintf("product:price:%s", product.ID.String())
            
            pipe := u.redisClient.Pool.Pipeline()

            pipe.Set(ctx, priceKey, product.StartingPrice, ttl)
            
            pipe.ZAdd(ctx, "hot_ranking", redis.Z{
                Score:  0,
                Member: product.ID.String(),
            })

            _, err = pipe.Exec(ctx)
            if err != nil {
                log.Printf("Redis initialization failed for product %s: %v", product.ID, err)
            }
        }
    }

    return nil
}

func (u *productUsecase) GetActiveAuctions(ctx context.Context, limit, offset int) ([]entity.Product, error) {
	return u.repo.GetActiveAuctions(ctx, limit, offset)
}

func (u *productUsecase) GetByID(ctx context.Context, id string) (*entity.Product, error) {
	return u.repo.GetByID(ctx, id)
}



