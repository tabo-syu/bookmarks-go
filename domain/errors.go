package domain

import "fmt"

type ValidationError struct {
	entity string
	field  string
	err    error
}

func NewValidationError(entity string, field string, err error) *ValidationError {
	return &ValidationError{
		entity: entity,
		field:  field,
		err:    err,
	}
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s.%s invalid value: %s", e.entity, e.field, e.err.Error())
}

func (e *ValidationError) Unwrap() error {
	return e.err
}
