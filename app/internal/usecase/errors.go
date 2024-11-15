package usecase

import "fmt"

type NotFountError struct {
	ID    string
	Model string
}

func (e *NotFountError) Error() string {
	return fmt.Sprintf("%s[%s] not found", e.Model, e.ID)
}

type UnexpectedError struct {
	Err error
}

func (e *UnexpectedError) Error() string {
	return "Unexpected error"
}
