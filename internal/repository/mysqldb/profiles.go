package mysqldb

import (
	"github.com/yendefrr/wg-admin/internal/models"
	"gorm.io/gorm"
)

type ProfileRepo struct {
	db *gorm.DB
}

func NewProfilesRepo(db *gorm.DB) *ProfileRepo {
	return &ProfileRepo{
		db: db,
	}
}

func (r *ProfileRepo) Create(form models.ProfileCreateForm) error {
	user := models.User{
		Username: form.Username,
	}

	profile := models.Profile{
		Username:  form.Username,
		Name:      form.Name,
		CreatedAt: "2023-02-19 00:00:00",
		UpdatedAt: "2023-02-19 00:00:00",
	}

	if res := r.db.Take(&user); res.RowsAffected != 1 {
		return gorm.ErrRecordNotFound
	}

	if res := r.db.Create(&profile); res.RowsAffected != 1 {
		return gorm.ErrInvalidTransaction
	}

	return nil
}

func (r *ProfileRepo) GetAll(isActive bool) ([]models.Profile, error) {
	var profiles []models.Profile

	if err := r.db.Where("is_active = ?", isActive).Find(&profiles).Error; err != nil {
		return nil, err
	}

	return profiles, nil
}

func (r *ProfileRepo) GetByID(id int) (*models.Profile, error) {
	p := &models.Profile{}

	return p, nil
}
