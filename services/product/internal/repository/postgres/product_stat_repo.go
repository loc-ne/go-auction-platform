package postgres

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/loc-ne/go-auction/services/product/internal/entity"
	"github.com/loc-ne/go-auction/services/product/internal/usecase"
)

type productStatRepo struct {
	db *pgxpool.Pool
}

func NewProductStatRepository(db *pgxpool.Pool) usecase.ProductStatRepository {
	return &productStatRepo{db: db}
}

func (r *productStatRepo) Create(ctx context.Context, productId uuid.UUID) error {
	sql := `INSERT INTO product_stats (product_id, bid_count, view_count, favorite_count) VALUES ($1, 0, 0, 0)`
	_, err := r.db.Exec(ctx, sql, productId)
	if err != nil {
		return err
	}
	return nil
}

func (r *productStatRepo) IncrementCounter(ctx context.Context, productID uuid.UUID, field string, amount int) (*entity.ProductStat, error) {
	var sql string
	switch field {
	case "bid_count":
		sql = `UPDATE product_stats SET bid_count = bid_count + $1, updated_at = NOW() WHERE product_id = $2 RETURNING product_id, bid_count, view_count, favorite_count, updated_at`
	case "view_count":
		sql = `UPDATE product_stats SET view_count = view_count + $1, updated_at = NOW() WHERE product_id = $2 RETURNING product_id, bid_count, view_count, favorite_count, updated_at`
	case "favorite_count":
		sql = `UPDATE product_stats SET favorite_count = favorite_count + $1, updated_at = NOW() WHERE product_id = $2 RETURNING product_id, bid_count, view_count, favorite_count, updated_at`
	default:
		return nil, errors.New("invalid counter field")
	}

	var productStat entity.ProductStat
	err := r.db.QueryRow(ctx, sql, amount, productID).Scan(
		&productStat.ProductID,
		&productStat.BidCount,
		&productStat.ViewCount,
		&productStat.FavoriteCount,
		&productStat.UpdatedAt,
	)
	
	if err != nil {
		return nil, err
	}

	return &productStat, nil
}

func (r *productStatRepo) GetStatByID(ctx context.Context, productID uuid.UUID) (*entity.ProductStat, error) {
	sql := `SELECT product_id, bid_count, view_count, favorite_count, updated_at FROM product_stats WHERE product_id = $1`
	var stat entity.ProductStat
	err := r.db.QueryRow(ctx, sql, productID).Scan(
		&stat.ProductID,
		&stat.BidCount,
		&stat.ViewCount,
		&stat.FavoriteCount,
		&stat.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &stat, nil
}
