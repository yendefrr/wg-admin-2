package mysqldb

import (
	"database/sql"
	"fmt"

	"github.com/yendefrr/wg-admin/internal/models"
)

type ProfileRepo struct {
	db *sql.DB
}

func NewProfilesRepo(db *sql.DB) *ProfileRepo {
	return &ProfileRepo{
		db: db,
	}
}

func (r *ProfileRepo) Create(form models.ProfileCreateForm) error {
	if err := r.db.QueryRow(fmt.Sprintf("SELECT `id` FROM `users` WHERE `username` = '%s'", form.Username)).Scan(); err != nil {
		return err
	}

	r.db.QueryRow(fmt.Sprintf(
		"INSERT INTO `profiles` (`username`, `name`) VALUES ('%s', '%s')",
		form.Username, form.Name))

	if err := r.db.QueryRow(fmt.Sprintf("SELECT `id` FROM `profiles` WHERE `username` = '%s' and `name` = '%s'", form.Username, form.Name)).Scan(); err != nil {
		return err
	}

	return nil
}

func (r *ProfileRepo) GetAll(isActive bool) ([]models.Profile, error) {
	res, err := r.db.Query("SELECT id, username, name FROM `profiles` WHERE is_active = ?", isActive)
	if err != nil {
		return nil, err
	}

	var profiles []models.Profile
	for res.Next() {
		var profile models.Profile
		err = res.Scan(
			&profile.ID,
			&profile.Username,
			&profile.Name,
		)
		if err != nil {
			return nil, err
		}

		profiles = append(profiles, profile)
	}

	return profiles, nil
}

func (r *ProfileRepo) GetByID(id int) (*models.Profile, error) {
	p := &models.Profile{}

	return p, nil
}
