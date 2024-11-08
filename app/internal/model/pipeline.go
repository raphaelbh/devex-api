package model

import (
	"encoding/json"
	"time"
)

type Pipeline struct {
	Model

	Name       string          `gorm:"not null;size:50"`
	Definition json.RawMessage `gorm:"not null;type:json"`

	Executions []PipelineExecution `gorm:"foreignKey:PipelineID"`
	Variables  []PipelineVariable  `gorm:"foreignKey:PipelineID"`
}

type PipelineExecution struct {
	Model
	StartingAt  time.Time
	FinishingAt time.Time

	TriggerType string          `gorm:"not null;size:20"` // [ MANUAL | PIPELINE | GIT | API ]
	Status      string          `gorm:"not null;size:20"` // [ WAITING_EXECUTION | IN_PROGRESS | CANCELED | FAIL | COMPLETED ]
	Input       json.RawMessage `gorm:"not null;type:json"`
	Logs        string          `gorm:"not null;type:text"`
	ElapsedTime int             `gorm:"not null"`

	PipelineID string `gorm:"not null"`
}

type PipelineVariable struct {
	Model

	Type  string `gorm:"not null;size:10"` // [ VARIABLE | SECRET ]
	Key   string `gorm:"not null;size:50"`
	Value string `gorm:"not null;size:200"`

	PipelineID string `gorm:"not null"`
}
