package redis

import (
	"context"
	"log"
	"os"
	"time"
	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	Pool *redis.Client
}

func NewRedisClient() (*RedisClient, error) {
	url := os.Getenv("REDIS_URL")
	
	opt, err := redis.ParseURL(url)
	if err != nil {
		return nil, err
	}

	opt.PoolSize = 20
	opt.MinIdleConns = 5
	opt.DialTimeout = 5 * time.Second

	rdb := redis.NewClient(opt)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	log.Println("Connected to Redis successfully")
	
	return &RedisClient{Pool: rdb}, nil
}

func (r *RedisClient) Publish(ctx context.Context, channel string, payload interface{}) error {
	return r.Pool.Publish(ctx, channel, payload).Err()
}

func (r *RedisClient) Subscribe(ctx context.Context, channel string) *redis.PubSub {
	return r.Pool.Subscribe(ctx, channel)
}

func (r *RedisClient) Get(ctx context.Context, key string) (string, error) {
	return r.Pool.Get(ctx, key).Result()
}

func (r *RedisClient) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	return r.Pool.HGetAll(ctx, key).Result()
}

func (r *RedisClient) PushAndTrimBid(ctx context.Context, productID string, bidJSON string, limit int64) error {
    key := "bid_history:" + productID
    pipe := r.Pool.Pipeline()
    pipe.LPush(ctx, key, bidJSON)
    pipe.LTrim(ctx, key, 0, limit-1)
    _, err := pipe.Exec(ctx)
    return err
}

// xử lý ttl
func (r *RedisClient) IncrViewerCount(ctx context.Context, roomID string) (int64, error) {
	return r.Pool.Incr(ctx, "viewer_count:" + roomID).Result()
}

func (r *RedisClient) DecrViewerCount(ctx context.Context, roomID string) (int64, error) {
	return r.Pool.Decr(ctx, "viewer_count:" + roomID).Result()
}

func (r *RedisClient) GetBidHistory(ctx context.Context, productID string, limit int64) ([]string, error) {
	return r.Pool.LRange(ctx, "bid_history:"+productID, 0, limit-1).Result()
}	