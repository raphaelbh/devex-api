package repository

import (
	"github.com/raphaelbh/devex-api/internal/model"
)

var (
	secretRepository = SecretRepository{}
)

func GetSecretRepository() SecretRepository {
	return secretRepository
}

type SecretRepository struct{}

func (r SecretRepository) Create(secret *model.Secret) error {
	return db.Create(secret).Error
}

func (r SecretRepository) FindByKey(key string) (*model.Secret, error) {
	var entity = model.Secret{}
	return &entity, db.Where("key = ?", key).Take(&entity).Error
}
