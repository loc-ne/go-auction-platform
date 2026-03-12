package redis

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/loc-ne/go-auction/services/product/internal/entity"
	"github.com/redis/go-redis/v9"
)

type RedisRepository interface {
	Publish(ctx context.Context, channel string, payload interface{}) error
	Subscribe(ctx context.Context, channel string) *redis.PubSub
	SetProductInitialState(ctx context.Context, product *entity.Product, ttl time.Duration) error
	UpdateHotRanking(ctx context.Context, productID string, score float64) error
	GetHotRanking(ctx context.Context, limit int64) ([]string, error)
	Close() error
}

type redisRepo struct {
	pool *redis.Client
}

func NewRedisRepository() (RedisRepository, error) {
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

	return &redisRepo{pool: rdb}, nil
}

func (r *redisRepo) Publish(ctx context.Context, channel string, payload interface{}) error {
	return r.pool.Publish(ctx, channel, payload).Err()
}

func (r *redisRepo) Subscribe(ctx context.Context, channel string) *redis.PubSub {
	return r.pool.Subscribe(ctx, channel)
}

func (r *redisRepo) SetProductInitialState(ctx context.Context, product *entity.Product, ttl time.Duration) error {
	priceKey := fmt.Sprintf("product:price:%s", product.ID.String())
	pipe := r.pool.Pipeline()
	fields := map[string]interface{}{
		"current_price": product.StartingPrice, 
		"bid_increment": product.BidIncrement,
		"seller_id":     product.SellerID.String(),
		"status":        "active",
	}

	pipe.HSet(ctx, priceKey, fields)
	pipe.Expire(ctx, priceKey, ttl)

	pipe.ZAdd(ctx, "hot_ranking", redis.Z{
		Score:  0,
		Member: product.ID.String(),
	})

	_, err := pipe.Exec(ctx)
	return err
}

func (r *redisRepo) UpdateHotRanking(ctx context.Context, productID string, score float64) error {
	return r.pool.ZAdd(ctx, "hot_ranking", redis.Z{
		Score:  score,
		Member: productID,
	}).Err()
}

func (r *redisRepo) GetHotRanking(ctx context.Context, limit int64) ([]string, error) {
	return r.pool.ZRevRange(ctx, "hot_ranking", 0, limit-1).Result()
}

func (r *redisRepo) Close() error {
	if r.pool != nil {
		return r.pool.Close()
	}
	return nil
}