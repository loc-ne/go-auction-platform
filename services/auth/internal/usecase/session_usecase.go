package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/loc-ne/go-auction/services/auth/internal/entity"
)
var (
    ErrSessionNotFound = errors.New("session not found")
    ErrSessionExpired = errors.New("session expired")
	ErrSessionBlocked = errors.New("session has been blocked")
)

type SessionRepository interface {
    Create(ctx context.Context, session *entity.Session) error
	Block(ctx context.Context, id uuid.UUID) error
	GetByID(ctx context.Context, id uuid.UUID) (*entity.Session, error)
}

type SessionUsecase interface {
	CreateSession(ctx context.Context, session *entity.Session) error
	BlockSession(ctx context.Context, id uuid.UUID) error
	GetSessionByID(ctx context.Context, id uuid.UUID) (*entity.Session, error)
}

type sessionUsecase struct {
    repo SessionRepository 
}

func NewSessionUsecase(r SessionRepository) SessionUsecase {
    return &sessionUsecase{repo: r}
}

func (u *sessionUsecase) CreateSession(ctx context.Context, session *entity.Session) error {
	session.ID = uuid.New()
	return u.repo.Create(ctx, session)
}

func (u *sessionUsecase) BlockSession(ctx context.Context, id uuid.UUID) error {
	return u.repo.Block(ctx, id)
}

func (u *sessionUsecase) GetSessionByID(ctx context.Context, id uuid.UUID) (*entity.Session, error) {
    session, err := u.repo.GetByID(ctx, id)
    if err != nil {
        return nil, err
    }
	if session == nil {
		return nil, ErrSessionNotFound
	}
    if session.IsBlocked {
        return nil, ErrSessionBlocked
    }
    if time.Now().After(session.ExpiresAt) {
        return nil, ErrSessionExpired
    }

    return session, nil
}