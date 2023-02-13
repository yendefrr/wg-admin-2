package service

import (
	"context"

	"github.com/yendefrr/wg-admin/internal/models"
	"github.com/yendefrr/wg-admin/internal/repository"
)

type ProfilesService struct {
	repo repository.Profiles
}

func NewProfilesService(repo repository.Profiles) *ProfilesService {
	return &ProfilesService{
		repo: repo,
	}
}

func (s *ProfilesService) Create(ctx context.Context, form models.ProfileCreateForm) error {

	return nil
}

func (s *ProfilesService) GetAllActive(ctx context.Context) ([]models.Profile, error) {
	profiles, err := s.repo.GetAll(true)
	if err != nil {
		return nil, err
	}

	return profiles, nil
}

func (s *ProfilesService) GetAllInActive(ctx context.Context) ([]models.Profile, error) {
	profiles, err := s.repo.GetAll(false)
	if err != nil {
		return nil, err
	}

	return profiles, nil
}
