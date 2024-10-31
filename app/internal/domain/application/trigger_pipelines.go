package application

import (
	"fmt"
	"log/slog"
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
	logger := slog.Default()
	logger.Info("TriggerPipelines")
	logger.Info(fmt.Sprintf("%+v", command))
	return nil
}
