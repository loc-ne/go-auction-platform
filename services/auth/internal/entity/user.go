package entity

import (
	"time"
	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `json:"id" db:"id"`
	Email        string    `json:"email" db:"email"`
	PasswordHash string    `json:"-" db:"password_hash"` 
	FullName     string    `json:"full_name" db:"full_name"`
	Role         string    `json:"role" db:"role"` // "admin", "bidder"
	Status       string    `json:"status" db:"status"` // "active", "blocked", "pending"
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}