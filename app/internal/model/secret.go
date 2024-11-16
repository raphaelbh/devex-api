package model

type Secret struct {
	Model

	Key   string `gorm:"not null;size:50"`
	Value string `gorm:"not null;size:250"`
}
