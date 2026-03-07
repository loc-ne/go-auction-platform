package usecase

import (
	"context"
	"github.com/loc-ne/go-auction/services/product/internal/entity"
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
}

func NewProductUsecase(r ProductRepository) ProductUsecase {
    return &productUsecase{
        repo:      r,
    }
}

func (u *productUsecase) Create(ctx context.Context, product *entity.Product) error {
	return u.repo.Create(ctx, product)
}

func (u *productUsecase) GetActiveAuctions(ctx context.Context, limit, offset int) ([]entity.Product, error) {
	return u.repo.GetActiveAuctions(ctx, limit, offset)
}

func (u *productUsecase) GetByID(ctx context.Context, id string) (*entity.Product, error) {
	return u.repo.GetByID(ctx, id)
}



