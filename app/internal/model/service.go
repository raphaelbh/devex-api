package model

import (
	"encoding/json"
	"time"
)

type Service struct {
	Model

	Name          string          `gorm:"not null;size:50"`
	RepositoryApp string          `gorm:"not null;size:100"`
	EnvVariables  json.RawMessage `gorm:"not null;type:jsonb"`

	TemplateID    string         `gorm:"not null"`
	ServiceEvents []ServiceEvent `gorm:"foreignKey:ServiceID"`
}

type ServiceEvent struct {
	Model
	TriggedAt time.Time `gorm:"not null"`

	Type string          `gorm:"not null;size:10"` // [ BUILD | DEPLOY ]
	Data json.RawMessage `gorm:"not null;type:jsonb"`

	ServiceID string `gorm:"not null"`
}
