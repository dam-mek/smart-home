package inmemory

import (
	"context"
	"errors"
	"homework/internal/domain"
	"sync"
)

var ErrUserNotFound = errors.New("user not found")

type UserRepository struct {
	users []*domain.User
	uMu   sync.Mutex
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) SaveUser(ctx context.Context, user *domain.User) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if user == nil {
		return errors.New("user is nil")
	}
	r.uMu.Lock()
	r.users = append(r.users, user)
	r.uMu.Unlock()
	return nil
}

func (r *UserRepository) GetUserByID(ctx context.Context, id int64) (*domain.User, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	for _, user := range r.users {
		if user.ID == id {
			return user, nil
		}
	}
	return nil, ErrUserNotFound
}
