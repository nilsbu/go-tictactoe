package util

import (
	"errors"
	"fmt"
)

// NewError returns an error with a specified text formatted in the Sprintf
// format.
func NewError(format string, args ...interface{}) error {
	// TODO remove and replace by Errorf
	return errors.New(fmt.Sprintf(format, args...))
}
