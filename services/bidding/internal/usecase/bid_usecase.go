package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/loc-ne/go-auction/services/bidding/internal/entity"
	"github.com/loc-ne/go-auction/services/bidding/internal/repository/redis"
)

type BidMessage struct {
	ProductID string  `json:"product_id"`
	Price     int64   `json:"price"`
}

type BidRepository interface {
    Create(ctx context.Context, bid *entity.Bid) error
}

type BidUsecase interface {
	CreateBid(ctx context.Context, bid *entity.Bid) error
	CheckRoomActive(ctx context.Context, productID string) (bool, error)
}

type bidUsecase struct {
    repo        BidRepository 
	redisClient *redis.RedisClient
}

func NewBidUsecase(r BidRepository, red *redis.RedisClient) BidUsecase {
    return &bidUsecase{
        repo:        r,
		redisClient: red,
    }
}

func (u *bidUsecase) CreateBid(ctx context.Context, bid *entity.Bid) error {
	err := u.repo.Create(ctx, bid)
	if err != nil {
		return err
	}

	channelName := "bid_created"
	msg := BidMessage{
	ProductID: bid.ProductID.String(),
	Price:     bid.Amount,
    }

	payload, err := json.Marshal(msg)
    if err != nil {
        return err
    }
	_ = u.redisClient.Publish(ctx, channelName, payload)

	return nil
}

func (uc *bidUsecase) CheckRoomActive(ctx context.Context, productID string) (bool, error) {
	priceKey := fmt.Sprintf("product:price:%s", productID)
	status, err := uc.redisClient.Get(ctx, priceKey)
	if err != nil {
		return false, err 
	}
	
	if status == "active" {
		return true, nil
	}
	return false, nil
}
