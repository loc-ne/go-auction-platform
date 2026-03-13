package worker

import (
	"context"
	"encoding/json"
	"log"

	"github.com/google/uuid"
	"github.com/loc-ne/go-auction/services/product/internal/repository/redis"
	"github.com/loc-ne/go-auction/services/product/internal/usecase"
)


type ProductStatWorker struct {
	redisClient        redis.RedisRepository
	productStatUsecase usecase.ProductStatUsecase
	productUsecase     usecase.ProductUsecase
}

func NewProductStatWorker(redisClient redis.RedisRepository, productStatUsecase usecase.ProductStatUsecase, productUsecase usecase.ProductUsecase) *ProductStatWorker {
	return &ProductStatWorker{
		redisClient:        redisClient,
		productStatUsecase: productStatUsecase,
		productUsecase:     productUsecase,
	}
}


func (w *ProductStatWorker) Start(ctx context.Context) {
    sub := w.redisClient.Subscribe(ctx, "bid_created")
    defer sub.Close()

    log.Println("Product Worker: Listening for bid events...")

    for msg := range sub.Channel() {
        go func(payload string) {
            var event struct {
                UserID string  `json:"user_id"`
                ProductID string  `json:"product_id"`
                Amount    int64   `json:"price"` 
            }

            if err := json.Unmarshal([]byte(payload), &event); err != nil {
                log.Printf("Failed to unmarshal: %v", err)
                return
            }

            pID, err := uuid.Parse(event.ProductID)
            if err != nil {
                log.Printf("Invalid UUID: %v", err)
                return
            }

            stats, err := w.productStatUsecase.IncrementCounter(context.Background(), pID, "bid_count", 1)
            if err != nil {
                log.Printf("Failed to increment bid count: %v", err)
                return
            }

            if err := w.productUsecase.UpdateCurrentPrice(context.Background(), pID, event.Amount); err != nil {
                log.Printf("Failed to update price: %v", err)
            }

            if err := w.productStatUsecase.RefreshHotRanking(context.Background(), stats); err != nil {
                log.Printf("Failed to refresh ranking: %v", err)
            }
            
        }(msg.Payload)
    }
}