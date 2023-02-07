package domain

import "fmt"

type ValidationError struct {
	Entity string
	Field  string
	Err    error
}

func NewValidationError(entity string, field string, err error) *ValidationError {
	return &ValidationError{
		Entity: entity,
		Field:  field,
		Err:    err,
	}
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s.%s invalid value: %s", e.Entity, e.Field, e.Err.Error())
}

func (e *ValidationError) Unwrap() error {
	return e.Err
}
