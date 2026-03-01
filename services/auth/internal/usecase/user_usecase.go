package usecase

import (
	"context"
	"errors"
	"github.com/loc-ne/go-auction/services/auth/internal/entity"
	"golang.org/x/crypto/bcrypt"
	"github.com/google/uuid"
	"github.com/loc-ne/go-auction/shared/pkg"
)
var (
    ErrUserAlreadyExists  = errors.New("user already exists")
    ErrInvalidCredentials = errors.New("login information is incorrect")
)

type UserRepository interface {
    Create(ctx context.Context, user *entity.User) error
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
}

type UserUsecase interface {
	Register(ctx context.Context, fullName, email, password string) error
	Login(ctx context.Context, email string, password string) (*entity.User, string, string, error)
}

type userUsecase struct {
    repo UserRepository 
	jwtSecret string
}

func NewUserUsecase(r UserRepository, jwtSecret string) UserUsecase {
    return &userUsecase{
        repo:      r,
        jwtSecret: jwtSecret, 
    }
}
func (u *userUsecase) Register(ctx context.Context, fullName, email, password string) error {
	user, err := u.repo.GetByEmail(ctx, email)
	if err != nil {
		return err
	}	
	if user != nil {
		return ErrUserAlreadyExists
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	newUser := &entity.User{
        ID:           uuid.New(),
        Email:        email,
        PasswordHash: string(passwordHash),
		FullName:     fullName,
    }
	return u.repo.Create(ctx, newUser)
}

func (u *userUsecase) Login(ctx context.Context, email string, password string) (*entity.User, string, string, error) {
	user, err := u.repo.GetByEmail(ctx, email)
	if err != nil {
        return nil, "", "", err
    }
	if user != nil {
		err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
		if err != nil {
			return nil, "", "", ErrInvalidCredentials
		}	
	} else {
		return nil, "", "", ErrInvalidCredentials
	}
	
	accessToken, refreshToken, err := pkg.GenerateTokens(ctx, user.ID.String(), user.Email, u.jwtSecret)
    if err != nil {
        return nil, "", "", err
    }
    
    return user, accessToken, refreshToken, nil
}