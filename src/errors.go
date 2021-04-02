package main

import "github.com/pkg/errors"

const errorSeparator = " >> "

// WrapError Improve an error with a message giving it more context
func WrapError(err error, message string) error {
	return errors.Wrap(err, message + errorSeparator)
}
