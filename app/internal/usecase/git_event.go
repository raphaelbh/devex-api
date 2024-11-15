package usecase

import (
	"fmt"
)

type GitEventCommand struct {
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

func GitEvent(command GitEventCommand) error {
	logger.Info(fmt.Sprintf("%+v", command))
	return nil
}
