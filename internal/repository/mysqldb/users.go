package mysqldb

import (
	"github.com/yendefrr/wg-admin/internal/models"
	"gorm.io/gorm"
)

type UsersRepo struct {
	db *gorm.DB
}

func NewUsersRepo(db *gorm.DB) *UsersRepo {
	return &UsersRepo{
		db: db,
	}
}

func (r *UsersRepo) Create(form models.UserCreateForm) error {
	user := models.User{
		Username: form.Username,
	}

	if res := r.db.Take(&user); res.RowsAffected == 1 {
		return gorm.ErrInvalidData
	}

	user.PasswordHash = form.Password //TODO: hash password
	user.Role = form.Role

	if res := r.db.Create(&user); res.RowsAffected != 1 {
		return gorm.ErrInvalidTransaction
	}

	return nil
}

func (r *UsersRepo) GetAll() ([]models.User, error) {

	return nil, nil
}

func (r *UsersRepo) GetByUsername(username string) (*models.User, error) {
	u := &models.User{}

	if res := r.db.Where("username = ?", username).Find(&u); res.RowsAffected != 1 {
		return nil, gorm.ErrRecordNotFound
	}

	return u, nil
}
