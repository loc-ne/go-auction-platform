package models

import (
	"time"
	"github.com/google/uuid"
)

type Bid struct {
	ID        uuid.UUID `json:"id" db:"id"`
	ProductID uuid.UUID `json:"product_id" db:"product_id"` 
	UserID    uuid.UUID `json:"user_id" db:"user_id"`
	Amount    int64     `json:"amount" db:"amount"` 
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}