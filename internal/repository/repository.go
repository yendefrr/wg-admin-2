package repository

import (
	"github.com/yendefrr/wg-admin/internal/models"
	"gorm.io/gorm"

	"github.com/yendefrr/wg-admin/internal/repository/mysqldb"
)

type Users interface {
	Create(form models.UserCreateForm) error
	GetAll() ([]models.User, error)
	GetByUsername(username string) (*models.User, error)
}

type Profiles interface {
	Create(form models.ProfileCreateForm) error
	GetAll(IsActive bool) ([]models.Profile, error)
	GetByID(id int) (*models.Profile, error)
}
type Repositories struct {
	Users    Users
	Profiles Profiles
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		Users:    mysqldb.NewUsersRepo(db),
		Profiles: mysqldb.NewProfilesRepo(db),
	}
}
