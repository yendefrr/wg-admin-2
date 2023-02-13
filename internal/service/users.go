package service

import (
	"context"

	"github.com/yendefrr/wg-admin/internal/models"
	"github.com/yendefrr/wg-admin/internal/repository"
)

type UsersService struct {
	repo repository.Users
}

func NewUsersService(repo repository.Users) *UsersService {
	return &UsersService{
		repo: repo,
	}
}

func (s *UsersService) Create(ctx context.Context, form models.UserCreateForm) error {

	return nil
}
