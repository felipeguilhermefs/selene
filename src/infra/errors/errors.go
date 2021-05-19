package errors

import "github.com/pkg/errors"

const (
	// errorSeparator is just a format suggar to separate each known error step
	errorSeparator = " >> "
)

// Wrap Improve an error with a message giving it more context
func Wrap(err error, message string) error {
	return errors.Wrap(err, message+errorSeparator)
}
