package repository

import (
	"database/sql"

	"github.com/yendefrr/wg-admin/internal/models"

	"github.com/yendefrr/wg-admin/internal/repository/mysqldb"
)

type Users interface {
	GetAll() ([]models.User, error)
}

type Profiles interface {
	GetAll(IsActive bool) ([]models.Profile, error)
	GetByID(id int) (models.Profile, error)
}
type Repositories struct {
	Users    Users
	Profiles Profiles
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		Users:    mysqldb.NewUsersRepo(db),
		Profiles: mysqldb.NewProfilesRepo(db),
	}
}
