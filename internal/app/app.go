package app

import (
	"database/sql"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/yendefrr/wg-admin/internal/config"
	"github.com/yendefrr/wg-admin/internal/repository"
	"github.com/yendefrr/wg-admin/internal/server"
	"github.com/yendefrr/wg-admin/internal/service"
	"github.com/yendefrr/wg-admin/internal/transport/http/handler"

	_ "github.com/go-sql-driver/mysql"
)

func Run(configPath string) {
	cfg, err := config.Init(configPath)
	if err != nil {
		log.Fatalf("error occured while init config: %s", err.Error())
		return
	}

	log.Info(cfg.MySQL.DatabaseName)

	db, err := newDB(cfg.MySQL)
	if err != nil {
		log.Fatalf("error occured while connecting database: %s", err.Error())
		return
	}

	repos := repository.NewRepositories(db)
	usersService := service.NewUsersService(repos.Users)
	profilesService := service.NewProfilesService(repos.Profiles)

	handler := handler.NewHandler(usersService, profilesService)

	srv := server.NewServer(cfg, handler.Init())

	srv.Run()

	log.Info("Server started")
}

func newDB(cfg config.MySQLConfig) (*sql.DB, error) {
	log.Info(cfg.Password)
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", cfg.User, cfg.Password, cfg.URI, cfg.DatabaseName))
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, err
}
