package service

import (
	"context"
	"project/internal/domain"
)

type UserServiceI interface {
	GetByID(ctx context.Context, id int) (*domain.User, error)
}
