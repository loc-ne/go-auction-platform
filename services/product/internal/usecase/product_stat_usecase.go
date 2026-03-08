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

type ProductStatRepository interface {
    Create(ctx context.Context, productId uuid.UUID) error
    IncrementCounter(ctx context.Context, productID uuid.UUID, field string, amount int) (*entity.ProductStat, error)
}

type ProductStatUsecase interface {
	Create(ctx context.Context, productId uuid.UUID) error
	IncrementCounter(ctx context.Context, productID uuid.UUID, field string, amount int) (*entity.ProductStat, error)
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

func (u *productUsecase) Create(ctx context.Context, productId uuid.UUID) error {
    return u.repo.Create(ctx, productId)
}

func (u *productUsecase) IncrementCounter(ctx context.Context, productID uuid.UUID, field string, amount int) (*entity.ProductStat, error) {
    return u.repo.IncrementCounter(ctx, productID, field, amount)
}




