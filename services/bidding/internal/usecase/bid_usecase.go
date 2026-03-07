package usecase

import (
	"context"
	"github.com/loc-ne/go-auction/services/bidding/internal/entity"
)

type BidRepository interface {
    Create(ctx context.Context, bid *entity.Bid) error
}

type BidUsecase interface {
	CreateBid(ctx context.Context, bid *entity.Bid) error
}

type bidUsecase struct {
    repo BidRepository 
}

func NewBidUsecase(r BidRepository) BidUsecase {
    return &bidUsecase{
        repo:      r,
    }
}

func (u *bidUsecase) CreateBid(ctx context.Context, bid *entity.Bid) error {
	return u.repo.Create(ctx, bid)
}
