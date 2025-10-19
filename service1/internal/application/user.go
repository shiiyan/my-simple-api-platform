package application

import (
	commonapplication "github.com/shiiyan/my-simple-api-platform/pkg/application"
	"github.com/shiiyan/my-simple-api-platform/service1/internal/domain"
)

type UserUsecase struct {
	Repo domain.UserRepository
}

func NewUserUsecase(repo domain.UserRepository) *UserUsecase {
	return &UserUsecase{
		Repo: repo,
	}
}

func (s *UserUsecase) GetUser(id string) (*domain.User, error) {
	user, err := s.Repo.GetByID(id)
	if err != nil {
		return nil, commonapplication.ErrUserNotFound
	}
	return user, nil
}
