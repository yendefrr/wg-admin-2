package service

import (
	"context"
	"database/sql"

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
	if err := s.repo.Create(form); err != nil {
		return err
	}

	return nil
}

func (s *UsersService) GetByUsername(ctx context.Context, username string) bool {
	if _, err := s.repo.GetByUsername(username); err == sql.ErrNoRows {
		return false
	}

	return true
}
