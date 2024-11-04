package model

import "encoding/json"

type Template struct {
	Model

	Name          string          `gorm:"not null;size:50"`
	RepositoryApp string          `gorm:"not null;size:100"`
	Building      json.RawMessage `gorm:"type:json"`

	Services []Service
}
