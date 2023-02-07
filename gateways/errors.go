package gateways

import "fmt"

type ValidationError struct {
	Entity string
	Err    error
}

func NewValidationError(entity string, err error) *ValidationError {
	return &ValidationError{
		Entity: entity,
		Err:    err,
	}
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s invalid input value: %s", e.Entity, e.Err.Error())
}

type MissingEntityError struct {
	Err error
}

func NewMissingEntityError(err error) *MissingEntityError {
	return &MissingEntityError{
		Err: err,
	}
}

func (e *MissingEntityError) Error() string {
	return fmt.Sprintf("entity not found: %s", e.Err.Error())
}

type PersistenceError struct {
	Err error
}

func NewPersistenceError(err error) *PersistenceError {
	return &PersistenceError{
		Err: err,
	}
}

func (e *PersistenceError) Error() string {
	return fmt.Sprintf("persistence failed: %s", e.Err.Error())
}
