package repository

import (
	"context"
	"myapp/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) Create(ctx context.Context, u models.User) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO users (name, email, age) 
		VALUES ($1, $2, $3)
	`, u.Name, u.Email, u.Age)
	return err
}

func (r *UserRepository) GetAll(ctx context.Context) ([]models.User, error) {
	rows, err := r.db.Query(ctx, `SELECT id, name, email, age FROM users`)
	if err != nil {
		return nil, err
	}

	users := []models.User{}
	for rows.Next() {
		var u models.User
		rows.Scan(&u.ID, &u.Name, &u.Email, &u.Age)
		users = append(users, u)
	}
	return users, nil
}
