package usecase

import (
	"context"
    "log"
	"time"

	"github.com/google/uuid"
	"github.com/loc-ne/go-auction/services/product/internal/entity"
	repoRedis "github.com/loc-ne/go-auction/services/product/internal/repository/redis"
)

const (
    WeightBid      = 50.0
    WeightFavorite = 20.0
    WeightView     = 1.0
)

type ProductStatRepository interface {
    Create(ctx context.Context, productId uuid.UUID) error
    IncrementCounter(ctx context.Context, productID uuid.UUID, field string, amount int) (*entity.ProductStat, error)
    GetStatByID(ctx context.Context, productID uuid.UUID) (*entity.ProductStat, error)
}

type ProductStatUsecase interface {
	Create(ctx context.Context, productId uuid.UUID) error
	IncrementCounter(ctx context.Context, productID uuid.UUID, field string, amount int) (*entity.ProductStat, error)
    RefreshHotRanking(ctx context.Context, stat *entity.ProductStat) error
	QueueView(id string)
	RefreshHotRankingByID(ctx context.Context, productID string) error
}

type productStatUsecase struct {
    repo ProductStatRepository
	redisClient repoRedis.RedisRepository 
	viewChannel chan string
}

func NewProductStatUsecase(r ProductStatRepository, red repoRedis.RedisRepository) ProductStatUsecase {
    uc := &productStatUsecase{
        repo:      r,
		redisClient: red,
		viewChannel: make(chan string, 10000),
    }
	go uc.startViewWorker()
	return uc
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

    err := u.redisClient.UpdateHotRanking(ctx, stat.ProductID.String(), score)

    if err != nil {
        log.Printf("Ranking Error: [Product %s] %v", stat.ProductID, err)
        return err
    }

    return nil
}

func (u *productStatUsecase) QueueView(id string) {
	select {
	case u.viewChannel <- id:
	default:
		log.Println("view queue full, drop job")
	}
}

func (u *productStatUsecase) startViewWorker() {
    batch := make(map[string]int)
    ticker := time.NewTicker(10 * time.Second)

    for {
        select {
        case id := <-u.viewChannel:
            batch[id]++
            if batch[id] >= 100 {
                u.flushViews(batch)
                batch = make(map[string]int) 
            }
        case <-ticker.C:
            if len(batch) > 0 {
                u.flushViews(batch)
                batch = make(map[string]int) 
            }
        }
    }
}

func (u *productStatUsecase) flushViews(batch map[string]int) {
    for id, count := range batch {
        pID, err := uuid.Parse(id)
		if err != nil {
			continue
		}
        stats, err := u.repo.IncrementCounter(context.Background(), pID, "view_count", count)
        if err == nil {
            _ = u.RefreshHotRanking(context.Background(), stats)
        }
    }
}

func (u *productStatUsecase) RefreshHotRankingByID(ctx context.Context, productID string) error {
    pID, err := uuid.Parse(productID)
    if err != nil {
        return err
    }
    
    stat, err := u.repo.GetStatByID(ctx, pID)
    if err != nil {
        log.Printf("Failed to get stat for ranking: %v", err)
        return err
    }
    
    return u.RefreshHotRanking(ctx, stat)
}
