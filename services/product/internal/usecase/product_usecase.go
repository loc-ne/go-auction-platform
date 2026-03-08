package usecase

import (
	"context"
    "time"
    "log"

	"github.com/google/uuid"
	"github.com/loc-ne/go-auction/services/product/internal/entity"
	repoRedis "github.com/loc-ne/go-auction/services/product/internal/repository/redis"
)

type ProductRepository interface {
    Create(ctx context.Context, product *entity.Product) error
    GetActiveAuctions(ctx context.Context, limit, offset int) ([]entity.Product, error)
    GetByID(ctx context.Context, id string) (*entity.Product, error)
	UpdateCurrentPrice(ctx context.Context, productID uuid.UUID, currentPrice int64) error
	HandleFavorite(ctx context.Context, userID, productID uuid.UUID) (bool, error)
}

type ProductUsecase interface {
	Create(ctx context.Context, product *entity.Product) error
	GetActiveAuctions(ctx context.Context, limit, offset int) ([]entity.Product, error)
	GetByID(ctx context.Context, id string) (*entity.Product, error)
	UpdateCurrentPrice(ctx context.Context, productID uuid.UUID, currentPrice int64) error
	HandleFavorite(ctx context.Context, userID, productID uuid.UUID) (bool, error)
}

type productUsecase struct {
    repo ProductRepository
	redisClient repoRedis.RedisRepository 
}

func NewProductUsecase(r ProductRepository, red repoRedis.RedisRepository) ProductUsecase {
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
            err = u.redisClient.SetProductInitialState(ctx, product.ID.String(), product.StartingPrice, ttl)
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

func (u *productUsecase) UpdateCurrentPrice(ctx context.Context, productID uuid.UUID, currentPrice int64) error {
	return u.repo.UpdateCurrentPrice(ctx, productID, currentPrice)
}

func (u *productUsecase) HandleFavorite(ctx context.Context, userID, productID uuid.UUID) (bool, error) {
	return u.repo.HandleFavorite(ctx, userID, productID)
}


