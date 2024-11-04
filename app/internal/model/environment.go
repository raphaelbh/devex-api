package model

import "encoding/json"

type Environment struct {
	Model

	Name       string          `gorm:"not null;size:50"`
	Deployment json.RawMessage `gorm:"type:json"`
}
