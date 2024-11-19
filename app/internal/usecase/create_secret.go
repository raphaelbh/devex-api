package usecase

import (
	"github.com/raphaelbh/devex-api/internal/commons/encryption"
	"github.com/raphaelbh/devex-api/internal/model"
)

type CreateSecretCommand struct {
	Key   string
	Value string
}

func CreateSecret(command CreateSecretCommand) error {
	encryptedValue, _ := encryption.Encrypt(command.Value)
	secretRepository.Create(&model.Secret{
		Key:   command.Key,
		Value: encryptedValue,
	})
	return nil
}
