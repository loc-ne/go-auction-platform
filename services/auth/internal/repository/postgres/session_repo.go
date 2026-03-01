package postgres

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/loc-ne/go-auction/services/auth/internal/entity"
)

type sessionRepo struct {
    db *pgxpool.Pool
}

func NewSessionRepository(db *pgxpool.Pool) *sessionRepo {
    return &sessionRepo{db: db}
}

func (r *sessionRepo) Create(ctx context.Context, s *entity.Session) error {
    sql := `
        INSERT INTO sessions (id, user_id, refresh_token, user_agent, client_ip, expires_at) 
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING created_at`
    
    err := r.db.QueryRow(ctx, sql, s.ID, s.UserID, s.RefreshToken, s.UserAgent, s.ClientIP, s.ExpiresAt).Scan(&s.CreatedAt)
    return err
}

func (r *sessionRepo) Block(ctx context.Context, id uuid.UUID) error {
    sql := `UPDATE sessions SET is_blocked = true WHERE id = $1`
    _, err := r.db.Exec(ctx, sql, id)
    return err
}

func (r *sessionRepo) GetByID(ctx context.Context, id uuid.UUID) (*entity.Session, error) {
    var s entity.Session
    sql := `SELECT id, user_id, refresh_token, is_blocked, expires_at FROM sessions WHERE id = $1`
    err := r.db.QueryRow(ctx, sql, id).Scan(&s.ID, &s.UserID, &s.RefreshToken, &s.IsBlocked, &s.ExpiresAt)
    if err != nil {
        if errors.Is(err, pgx.ErrNoRows) {
            return nil, nil
        }
        return nil, err
    }
    return &s, nil
}