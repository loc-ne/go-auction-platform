package postgres

import (
	"context"
	"github.com/loc-ne/go-auction/services/bidding/internal/entity"
    "github.com/loc-ne/go-auction/services/bidding/internal/usecase" 
	"github.com/jackc/pgx/v5/pgxpool"
)	

type bidRepo struct {
    db *pgxpool.Pool
}

func NewBidRepository(db *pgxpool.Pool) usecase.BidRepository {
    return &bidRepo{db: db}	
}

func (r *bidRepo) Create(ctx context.Context, b *entity.Bid) error { 
	sql := `INSERT INTO bids (product_id, user_id, amount) VALUES ($1, $2, $3) RETURNING id, created_at`
	err := r.db.QueryRow(ctx, sql, b.ProductID, b.UserID, b.Amount).Scan(&b.ID, &b.CreatedAt)
	if err != nil {
		return err
	}
    return nil 
}

