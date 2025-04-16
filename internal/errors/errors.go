package errors

import "fmt"

func NewErrTagger(tag string) func(error) error {
	return func(err error) error {
		return fmt.Errorf("%s: %w", tag, err)
	}
}
