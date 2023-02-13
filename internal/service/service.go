package service

import (
	"context"

	"github.com/yendefrr/wg-admin/internal/models"
)

type Users interface {
	Create(ctx context.Context, form models.UserCreateForm) error
}

type Profiles interface {
	Create(ctx context.Context, form models.ProfileCreateForm) error
	GetAllActive(ctx context.Context) ([]models.Profile, error)
	GetAllInActive(ctx context.Context) ([]models.Profile, error)
}
