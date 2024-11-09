package usecase

import (
	"fmt"
)

type TriggerPipelinesCommand struct {
	Branch string
	Pusher struct {
		Name  string
		Email string
	}
	Repository struct {
		Name string
		Url  string
	}
}

func TriggerPipelines(command TriggerPipelinesCommand) error {
	logger.Info("TriggerPipelines")
	logger.Info(fmt.Sprintf("%+v", command))
	return nil
}
