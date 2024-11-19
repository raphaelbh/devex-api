package repository

import (
	"log/slog"

	"github.com/raphaelbh/devex-api/internal/commons/config"
	"github.com/raphaelbh/devex-api/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	conf = config.Default()
)

func init() {
	conn, err := gorm.Open(postgres.Open(conf.Database.Url), &gorm.Config{})
	if err != nil {
		logger.Error("Unable to connect to database", slog.String("error", err.Error()))
		panic("Unable to connect to database: " + err.Error())
	}
	db = conn
	logger.Info("Database connected")

	if err := conn.AutoMigrate(
		&model.Environment{},
		&model.Secret{},
		&model.Template{},
		&model.Service{},
		&model.ServiceEvent{},
		&model.Pipeline{},
		&model.PipelineExecution{},
		&model.PipelineVariable{},
	); err != nil {
		logger.Error("Unable to migrate database tables", slog.String("error", err.Error()))
		panic("Unable to migrate database tables: " + err.Error())
	}

	logger.Info("Database migration completed")
}
