package repository

import (
	"log/slog"
	"time"

	"github.com/raphaelbh/devex-api/internal/config"
	"github.com/raphaelbh/devex-api/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger = slog.Default()
	conf   = config.Default()
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
		&model.Template{},
		&model.Service{},
		&model.ServiceEvent{},
		&model.Environment{},
		&model.Pipeline{},
		&model.PipelineExecution{},
		&model.PipelineVariable{},
	); err != nil {
		logger.Error("Unable to migrate database tables", slog.String("error", err.Error()))
		panic("Unable to migrate database tables: " + err.Error())
	}

	logger.Info("Database migration completed")
}

type Model struct {
	ID        string `gorm:"primaryKey;size:27"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
