package mysqldb

import (
	"database/sql"

	"github.com/yendefrr/wg-admin/internal/models"
)

type UsersRepo struct {
	db *sql.DB
}

func NewUsersRepo(db *sql.DB) *UsersRepo {
	return &UsersRepo{
		db: db,
	}
}

func (r *UsersRepo) GetAll() ([]models.User, error) {

	return nil, nil
}
