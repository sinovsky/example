package service

import (
	"context"
	"project/internal/domain"
	"project/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (us *UserService) GetByID(ctx context.Context, id int) (*domain.User, error) {
	return us.repo.GetByID(ctx, id)
}
