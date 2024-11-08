package model

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        string         `gorm:"primaryKey;size:27"`
	CreatedAt time.Time      `gorm:"not null"`
	UpdatedAt time.Time      `gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
