package repository

import (
	"github.com/raphaelbh/devex-api/internal/model"
)

type PipelineRepository struct{}

func (r PipelineRepository) Create(model *model.Pipeline) error {
	return db.Create(model).Error
}

func (r PipelineRepository) FindById(id string) (*model.Pipeline, error) {
	var entity = model.Pipeline{}
	return &entity, db.Where("id = ?", id).Take(&entity).Error
}
