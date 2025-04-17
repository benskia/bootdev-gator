package gatorerrs

import (
	"fmt"
)

func NewErrWrapper(tag string) func(string, error) error {
	return func(desc string, err error) error {
		return fmt.Errorf("%s: %s\n\t%w", tag, desc, err)
	}
}
