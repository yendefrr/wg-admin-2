package tests

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"

	"github.com/yendefrr/wg-admin/internal/repository/mysqldb"
)

func TestProfileRepo_GetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("ошибка при создании mock-объекта: %v", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "username", "name"}).
		AddRow(1, "user1", "profile1").
		AddRow(2, "user2", "profile2")

	mock.ExpectQuery("SELECT id, username, name FROM `profiles` WHERE is_active = ?").
		WithArgs(true).
		WillReturnRows(rows)

	repo := mysqldb.NewProfilesRepo(db)
	profiles, err := repo.GetAll(true)
	if err != nil {
		t.Fatalf("ошибка при вызове GetAll(): %v", err)
	}

	if len(profiles) != 2 {
		t.Errorf("ожидалось 2 профиля, получено %d", len(profiles))
	}

	if profiles[0].ID != 1 || profiles[0].Username != "user1" || profiles[0].Name != "profile1" {
		t.Errorf("некорректный первый профиль: %v", profiles[0])
	}

	if profiles[1].ID != 2 || profiles[1].Username != "user2" || profiles[1].Name != "profile2" {
		t.Errorf("некорректный второй профиль: %v", profiles[1])
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("ожидаемые вызовы не были выполнены: %v", err)
	}
}
