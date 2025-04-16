package errors

import (
	"errors"
	"fmt"
)

var ErrLoginArgs error = errors.New("expected 'login <username>'")
var ErrNotFound error = errors.New("not found")

func NewErrTagger(tag string) func(error) error {
	return func(err error) error {
		return fmt.Errorf("%s: %w", tag, err)
	}
}
