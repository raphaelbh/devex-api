package model

import (
	"time"

	"github.com/segmentio/ksuid"
	"gorm.io/gorm"
)

type Model struct {
	ID        string         `gorm:"primaryKey;size:27"`
	CreatedAt time.Time      `gorm:"not null"`
	UpdatedAt time.Time      `gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (m *Model) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ID == "" {
		m.ID = ksuid.New().String()
	}
	return
}
