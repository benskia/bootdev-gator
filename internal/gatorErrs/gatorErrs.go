package gatorerrs

import (
	"fmt"
)

func NewErrWrapper(tag string) func(error) error {
	return func(err error) error {
		return fmt.Errorf("%s: %w", tag, err)
	}
}
