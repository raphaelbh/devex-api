package usecase

import (
	"errors"

	"github.com/raphaelbh/devex-api/internal/model"
	"gorm.io/gorm"
)

func GetEnvironment(id string) (*model.Environment, error) {
	var env, err = envRepository.Find(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &NotFountError{
				ID:    id,
				Model: "Environment",
			}
		}
		return nil, &UnexpectedError{}
	}

	return env, nil
}
