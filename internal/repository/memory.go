package repository

import (
	"context"
	"errors"
	"project/internal/domain"
	"sync"
	"time"
)

type MemoryUserRepository struct {
	mu    sync.RWMutex
	users map[int]*domain.User
	seq   int
}

func NewMemoryUserRepository() *MemoryUserRepository {
	data := make(map[int]*domain.User)
	fakeUser := domain.User{
		ID:       1,
		Username: "Fake User",
	}

	data[1] = &fakeUser
	return &MemoryUserRepository{
		users: data,
		seq:   0,
	}
}

func (ur *MemoryUserRepository) GetByID(ctx context.Context, id int) (*domain.User, error) {
	ur.mu.RLock()
	defer ur.mu.RUnlock()

	user, ok := ur.users[id]
	if !ok {
		return nil, errors.New("пользователь с таким id не найден")
	}
	return user, nil
}

func (ur *MemoryUserRepository) List(ctx context.Context) ([]*domain.User, error) {
	ur.mu.RLock()
	defer ur.mu.RUnlock()

	users := make([]*domain.User, 0, len(ur.users))
	for _, user := range ur.users {
		users = append(users, user)
	}
	return users, nil
}

func (ur *MemoryUserRepository) Create(ctx context.Context, user *domain.User) error {
	ur.mu.RLock()
	defer ur.mu.RUnlock()

	ur.seq++
	user.ID = ur.seq
	ur.users[user.ID] = user
	return nil
}

func (ur *MemoryUserRepository) Update(ctx context.Context, user *domain.User) error {
	ur.mu.RLock()
	defer ur.mu.RUnlock()

	if _, ok := ur.users[user.ID]; !ok {
		return errors.New("пользователь с таким id не найден")
	}

	user.UpdatedAt = time.Now()
	ur.users[user.ID] = user
	return nil

}

func (ur *MemoryUserRepository) Delete(ctx context.Context, id int) error {
	ur.mu.RLock()
	defer ur.mu.RUnlock()

	if _, ok := ur.users[id]; !ok {
		return errors.New("пользователь с таким id не найден")
	}

	delete(ur.users, id)
	return nil
}
