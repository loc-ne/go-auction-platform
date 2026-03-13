package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/loc-ne/go-auction/services/bidding/internal/entity"
	"github.com/loc-ne/go-auction/services/bidding/internal/repository/redis"
)

var (
	ErrProductNotFound  = errors.New("product not found or auction closed")
	ErrBidPriceTooLow   = errors.New("bid price must be greater than current price + bid increment")
	ErrSellerCannotBid  = errors.New("seller cannot bid on their own product")
	ErrAuctionNotActive = errors.New("auction is not currently active")
)

type BidMessage struct {
	UserID string  `json:"user_id"`
	ProductID string  `json:"product_id"`
	Price     int64   `json:"price"`
}

type BidHistory struct {
	Price     int64   `json:"price"`
	BidderName string  `json:"bidder_name"`
	CreatedAt time.Time `json:"created_at"`
}

type BidRepository interface {
    Create(ctx context.Context, bid *entity.Bid) error
}

type BidUsecase interface {
	CreateBid(ctx context.Context, bid *entity.Bid, bidderName string) error
	CheckRoomActive(ctx context.Context, productID string) (bool, error)
	ValidateBid(ctx context.Context, productID string, price int64, userID uuid.UUID) error
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

func (u *bidUsecase) CreateBid(ctx context.Context, bid *entity.Bid, bidderName string) error {
	err := u.repo.Create(ctx, bid)
	if err != nil {
		return err
	}

	channelName := "bid_created"
	msg := BidMessage{
		UserID: bid.UserID.String(),
		ProductID: bid.ProductID.String(),
		Price:     bid.Amount,
    }

	payload, err := json.Marshal(msg)
    if err != nil {
        return err
    }
	_ = u.redisClient.Publish(ctx, channelName, payload)

	bidHistory := BidHistory{
		Price:     bid.Amount,
		BidderName:    bidderName,
		CreatedAt: bid.CreatedAt,
	}
	
	bidJSON, err := json.Marshal(bidHistory)
	if err != nil {
		return err
	}
	_ = u.redisClient.PushAndTrimBid(ctx, bid.ProductID.String(), string(bidJSON), 10)

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

func (uc *bidUsecase) ValidateBid(ctx context.Context, productID string, price int64, userID uuid.UUID) error {
	priceKey := fmt.Sprintf("product:price:%s", productID)
	
	redisData, err := uc.redisClient.HGetAll(ctx, priceKey)
	if err != nil {
		return err
	}
	
	if len(redisData) == 0 {
		return ErrProductNotFound
	}
	
	if redisData["status"] != "active" {
		return ErrAuctionNotActive
	}
	
	currentPrice, _ := strconv.ParseInt(redisData["current_price"], 10, 64)
	bidIncrement, _ := strconv.ParseInt(redisData["bid_increment"], 10, 64)
	sellerID := redisData["seller_id"]

	if price < currentPrice + bidIncrement {
		return ErrBidPriceTooLow
	}
	
	if sellerID == userID.String() {
		return ErrSellerCannotBid
	}
	
	return nil
}
