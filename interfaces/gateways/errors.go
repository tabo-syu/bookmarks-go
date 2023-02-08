package gateways

import "fmt"

type MissingEntityError struct {
	err error
}

func NewMissingEntityError(err error) *MissingEntityError {
	return &MissingEntityError{
		err: err,
	}
}

func (e *MissingEntityError) Error() string {
	return fmt.Sprintf("entity not found: %s", e.err.Error())
}

type PersistenceError struct {
	err error
}

func NewPersistenceError(err error) *PersistenceError {
	return &PersistenceError{
		err: err,
	}
}

func (e *PersistenceError) Error() string {
	return fmt.Sprintf("persistence failed: %s", e.err.Error())
}
