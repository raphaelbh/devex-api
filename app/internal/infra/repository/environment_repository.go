package repository

import (
	"github.com/raphaelbh/devex-api/internal/model"
)

var (
	environmentRepository = EnvironmentRepository{}
)

func GetEnvironmentRepository() EnvironmentRepository {
	return environmentRepository
}

type EnvironmentRepository struct{}

func (r EnvironmentRepository) FindAll() (*[]model.Environment, error) {
	var entities = []model.Environment{}
	return &entities, db.Find(&entities).Error
}

func (r EnvironmentRepository) Find(id string) (*model.Environment, error) {
	var entity = model.Environment{}
	return &entity, db.Where("id = ?", id).Take(&entity).Error
}
