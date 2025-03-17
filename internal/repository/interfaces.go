package repository

import (
	"context"
	"project/internal/domain"
)

type UserRepository interface {
	GetByID(ctx context.Context, id int) (*domain.User, error)
	List(ctx context.Context) ([]*domain.User, error)
	Create(ctx context.Context, user *domain.User) error
	Update(ctx context.Context, user *domain.User) error
	Delete(ctx context.Context, id int) error
}
