package postgres

import (
	"context"
	"github.com/loc-ne/go-auction/services/product/internal/entity"
    "github.com/loc-ne/go-auction/services/product/internal/usecase" 
	"github.com/jackc/pgx/v5/pgxpool"
)	

type productRepo struct {
    db *pgxpool.Pool
}

func NewProductRepository(db *pgxpool.Pool) usecase.ProductRepository {
    return &productRepo{db: db}	
}

func (r *productRepo) Create(ctx context.Context, product *entity.Product) error {
	sql := `INSERT INTO products (seller_id, name, description, starting_price, current_price, bid_increment, status, image_urls, start_at, end_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id, created_at, updated_at`
	err := r.db.QueryRow(ctx, sql, product.SellerID, product.Name, product.Description, product.StartingPrice, product.CurrentPrice, product.BidIncrement, product.Status, product.ImageURLs, product.StartAt, product.EndAt).Scan(&product.ID, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (r *productRepo) GetActiveAuctions(ctx context.Context, limit, offset int) ([]entity.Product, error) {
	sql := `SELECT id, seller_id, name, description, starting_price, current_price, bid_increment, status, image_urls, start_at, end_at, winner_id, created_at, updated_at FROM products WHERE status = 1 ORDER BY start_at DESC LIMIT $1 OFFSET $2`
	rows, err := r.db.Query(ctx, sql, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	products := make([]entity.Product, 0)
	for rows.Next() {
		var product entity.Product
		err := rows.Scan(&product.ID, &product.SellerID, &product.Name, &product.Description, &product.StartingPrice, &product.CurrentPrice, &product.BidIncrement, &product.Status, &product.ImageURLs, &product.StartAt, &product.EndAt, &product.WinnerID, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}

func (r *productRepo) GetByID(ctx context.Context, id string) (*entity.Product, error) {
	sql := `SELECT id, seller_id, name, description, starting_price, current_price, bid_increment, status, image_urls, start_at, end_at, winner_id, created_at, updated_at FROM products WHERE id = $1`
	row := r.db.QueryRow(ctx, sql, id)
	var product entity.Product
	err := row.Scan(&product.ID, &product.SellerID, &product.Name, &product.Description, &product.StartingPrice, &product.CurrentPrice, &product.BidIncrement, &product.Status, &product.ImageURLs, &product.StartAt, &product.EndAt, &product.WinnerID, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

