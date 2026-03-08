package entity

import (
	"time"
	"github.com/google/uuid"
)

type ProductStat struct {
	ProductID      uuid.UUID    `json:"product_id" db:"product_id"`
	BidCount       int       `json:"bid_count" db:"bid_count"`
	ViewCount      int       `json:"view_count" db:"view_count"`
	FavoriteCount  int       `json:"favorite_count" db:"favorite_count"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}