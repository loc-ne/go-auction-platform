package entity

import (
	"time"
	"github.com/google/uuid"
)

type Product struct {
	ID           uuid.UUID `json:"id" db:"id"`
	SellerID     uuid.UUID `json:"seller_id" db:"seller_id"`
	Name         string    `json:"name" db:"name"`
	Description  string    `json:"description" db:"description"`
	StartingPrice int64     `json:"starting_price" db:"starting_price"`
	CurrentPrice int64     `json:"current_price" db:"current_price"`
	BidIncrement int64     `json:"bid_increment" db:"bid_increment"`
	Status       int       `json:"status" db:"status"` // 0: Draft, 1: Active, 2: Completed, 3: Canceled
	ImageURLs    []string  `json:"image_urls" db:"image_urls"`
	StartAt      time.Time `json:"start_at" db:"start_at"`
	EndAt        time.Time `json:"end_at" db:"end_at"`
	WinnerID     uuid.UUID `json:"winner_id" db:"winner_id"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}