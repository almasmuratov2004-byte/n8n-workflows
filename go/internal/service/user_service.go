package service

import (
	"context"
	"myapp/internal/models"
	"myapp/internal/repository"

	"github.com/go-playground/validator/v10"
)

type UserService struct {
	repo      *repository.UserRepository
	validator *validator.Validate
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo:      repo,
		validator: validator.New(),
	}
}

func (s *UserService) Create(ctx context.Context, u models.User) error {
	if err := s.validator.Struct(u); err != nil {
		return err
	}
	return s.repo.Create(ctx, u)
}

func (s *UserService) GetAll(ctx context.Context) ([]models.User, error) {
	return s.repo.GetAll(ctx)
}
