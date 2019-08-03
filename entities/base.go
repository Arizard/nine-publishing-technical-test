package entities

import (
	"errors"
)

type repositoryError struct {
	err string
}

func (e *repositoryError) Error() string {
	return fmt.Sprintf("Repository error: %s", e.err)
}