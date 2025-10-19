package db

import (
	commondomain "github.com/shiiyan/my-simple-api-platform/pkg/domain"
	"github.com/shiiyan/my-simple-api-platform/service1/internal/domain"
)

type InMemoryUserRepo struct {
	users map[string]*domain.User
}

func NewInMemoryUserRepo() *InMemoryUserRepo {
	return &InMemoryUserRepo{
		users: map[string]*domain.User{
			"u1": {ID: "u1", Name: "Alice"},
		},
	}
}

func (r *InMemoryUserRepo) GetByID(id string) (*domain.User, error) {
	user, ok := r.users[id]
	if !ok {
		return nil, commondomain.ErrUserNotFound
	}
	return user, nil
}
