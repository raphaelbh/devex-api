package usecase

import (
	"github.com/raphaelbh/devex-api/internal/infra/repository"
	"github.com/raphaelbh/devex-api/internal/model"
)

var (
	envRepository = repository.GetEnvironmentRepository()
)

func GetAllEnvironments() (*[]model.Environment, error) {
	var environments, err = envRepository.FindAll()
	if err != nil {
		return nil, &UnexpectedError{}
	}

	return environments, nil
}
