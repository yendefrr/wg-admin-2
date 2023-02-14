package mysqldb

import (
	"database/sql"
	"fmt"

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

func (r *UsersRepo) Create(form models.UserCreateForm) error {
	r.db.QueryRow(fmt.Sprintf(
		"INSERT INTO users (username, role) VALUES ('%s', '%s')", form.Username, form.Role))

	return r.db.QueryRow(fmt.Sprintf("SELECT `id` FROM `users` WHERE `username` = '%s'", form.Username)).Scan()
}

func (r *UsersRepo) GetAll() ([]models.User, error) {

	return nil, nil
}

func (r *UsersRepo) GetByUsername(username string) (*models.User, error) {
	u := &models.User{
		Username: username,
	}
	if err := r.db.QueryRow(
		fmt.Sprintf("SELECT `id`, `role` FROM `users` WHERE `username` = '%s'", username)).Scan(
		&u.ID,
		&u.Role,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}

		return nil, err
	}

	return u, nil
}
