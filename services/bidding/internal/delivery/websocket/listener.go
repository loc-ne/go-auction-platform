package websocket

import (
	"context"
	"encoding/json"
	"log"

	"github.com/loc-ne/go-auction/services/bidding/internal/repository/redis"
)

func StartRedisListener(ctx context.Context, hub *Hub, redisClient *redis.RedisClient) {
	sub := redisClient.Subscribe(ctx, "bid_created")
    defer sub.Close()

    log.Println("Bid Listener: Listening for bid events...")

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

            msg := Message{
                RoomID:  event.ProductID,
                UserID:  event.UserID,
                Action:  "bid_created",
                Payload: event.Amount,
            }

            hub.Broadcast(msg)

            
        }(msg.Payload)
    }
}