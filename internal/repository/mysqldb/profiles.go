package mysqldb

import (
	"database/sql"

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

func (r *ProfileRepo) GetByID(id int) (models.Profile, error) {
	var profile models.Profile

	return profile, nil
}
