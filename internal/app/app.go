package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/yendefrr/wg-admin/internal/config"
	"github.com/yendefrr/wg-admin/internal/repository"
	"github.com/yendefrr/wg-admin/internal/server"
	"github.com/yendefrr/wg-admin/internal/service"
	v1 "github.com/yendefrr/wg-admin/internal/transport/http/api/v1"
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
	api := v1.NewHandler(usersService, profilesService)

	srv := server.NewServer(cfg, handler.Init())
	cfg.HTTP.Port = "8080"
	srvAPI := server.NewServer(cfg, api.Init())

	go srv.Run()
	go srvAPI.Run()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	log.Info("Server started")
}

func newDB(cfg config.MySQLConfig) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@%s/%s", cfg.User, cfg.Password, cfg.URI, cfg.DatabaseName)), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, err
}
