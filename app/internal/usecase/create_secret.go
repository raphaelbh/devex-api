package usecase

import (
	"github.com/raphaelbh/devex-api/internal/commons/encryption"
	"github.com/raphaelbh/devex-api/internal/infra/repository"
	"github.com/raphaelbh/devex-api/internal/model"
)

var (
	secretRepository = repository.GetSecretRepository()
)

type CreateSecretCommand struct {
	Key   string
	Value string
}

func (c CreateSecretCommand) toModel() *model.Secret {
	encryptedValue, _ := encryption.Encrypt(c.Value)
	return &model.Secret{
		Key:   c.Key,
		Value: encryptedValue,
	}
}

func CreateSecret(command CreateSecretCommand) error {
	secretRepository.Create(command.toModel())
	return nil
}
