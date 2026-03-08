package usecase

import (
	"context"
    "log"

	"github.com/google/uuid"
	"github.com/loc-ne/go-auction/services/product/internal/entity"
	repoRedis "github.com/loc-ne/go-auction/services/product/internal/repository/redis"
    "github.com/redis/go-redis/v9"
)

const (
    WeightBid      = 50.0
    WeightFavorite = 20.0
    WeightView     = 1.0
)

type ProductStatRepository interface {
    Create(ctx context.Context, productId uuid.UUID) error
    IncrementCounter(ctx context.Context, productID uuid.UUID, field string, amount int) (*entity.ProductStat, error)
}

type ProductStatUsecase interface {
	Create(ctx context.Context, productId uuid.UUID) error
	IncrementCounter(ctx context.Context, productID uuid.UUID, field string, amount int) (*entity.ProductStat, error)
    RefreshHotRanking(ctx context.Context, stat *entity.ProductStat) error
}

type productStatUsecase struct {
    repo ProductStatRepository
	redisClient *repoRedis.RedisClient 
}

func NewProductStatUsecase(r ProductStatRepository, red *repoRedis.RedisClient) ProductStatUsecase {
    return &productStatUsecase{
        repo:      r,
		redisClient: red,
    }
}

func (u *productStatUsecase) Create(ctx context.Context, productId uuid.UUID) error {
    return u.repo.Create(ctx, productId)
}

func (u *productStatUsecase) IncrementCounter(ctx context.Context, productID uuid.UUID, field string, amount int) (*entity.ProductStat, error) {
    return u.repo.IncrementCounter(ctx, productID, field, amount)
}

func (u *productStatUsecase) RefreshHotRanking(ctx context.Context, stat *entity.ProductStat) error {
    score := (float64(stat.BidCount) * 50.0) + 
             (float64(stat.FavoriteCount) * 20.0) + 
             (float64(stat.ViewCount) * 1.0)

    err := u.redisClient.Pool.ZAdd(ctx, "hot_ranking", redis.Z{
        Score:  score,
        Member: stat.ProductID.String(),
    }).Err()

    if err != nil {
        log.Printf("Ranking Error: [Product %s] %v", stat.ProductID, err)
        return err
    }

    return nil
}


