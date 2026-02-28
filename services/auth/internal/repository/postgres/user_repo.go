package postgres

import (
	"context"
	"github.com/loc-ne/go-auction/services/auth/internal/entity"
    "github.com/loc-ne/go-auction/services/auth/internal/usecase" 
	"github.com/jackc/pgx/v5/pgxpool"
)	

type userRepo struct {
    db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) usecase.UserRepository {
    return &userRepo{db: db}	
}

func (r *userRepo) Create(u *entity.User) error { 
	sql := `INSERT INTO users (email, password_hash, full_name) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(context.Background(), sql, u.Email, u.PasswordHash, u.FullName)
	if err != nil {
		return err
	}
    return nil 
}

func (r *userRepo) GetByEmail(email string) (*entity.User, error) {
	sql := `SELECT * FROM users WHERE email = $1`
	row := r.db.QueryRow(context.Background(), sql, email)
	var user entity.User
	err := row.Scan(&user.ID, &user.Email, &user.PasswordHash, &user.FullName, &user.Role, &user.Status, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, nil
	}
	return &user, nil
}