package errors

import "github.com/pkg/errors"

const errorSeparator = " >> "

// Wrap Improve an error with a message giving it more context
func Wrap(err error, message string) error {
	return errors.Wrap(err, message+errorSeparator)
}
