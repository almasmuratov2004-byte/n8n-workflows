package di

import (
	"myapp/internal/db"
	"myapp/internal/repository"
	"myapp/internal/service"
)

type Container struct {
	UserService *service.UserService
}

func NewContainer(dburl string) (*Container, error) {
	pg, err := db.Connect(dburl)
	if err != nil {
		return nil, err
	}

	userRepo := repository.NewUserRepository(pg)
	userService := service.NewUserService(userRepo)

	return &Container{
		UserService: userService,
	}, nil
}
