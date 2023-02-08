package domain

import "fmt"

type ValidationError struct {
	entity string
	err    error
}

func NewValidationError(entity string, err error) *ValidationError {
	return &ValidationError{
		entity: entity,
		err:    err,
	}
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s invalid value: %s", e.entity, e.err.Error())
}

func (e *ValidationError) Unwrap() error {
	return e.err
}
