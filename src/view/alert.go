package view

import (
	"log"

	"github.com/felipeguilhermefs/selene/core/auth"
	"github.com/felipeguilhermefs/selene/core/bookshelf"
)

const (
	// AlertLvError error alert level
	AlertLvError = "danger"
	// AlertLvWarning  warn alert level
	AlertLvWarning = "warning"
	// AlertLvInfo  info alert level
	AlertLvInfo = "info"
	// AlertLvSuccess  success alert level
	AlertLvSuccess = "success"

	defaultErrorMessage = "Something whent wrong. Please contact us if this error persists"
)

// Alert is used to render alert messages in templates
type Alert struct {
	Level   string
	Message string
}

func newErrorAlert(err error) *Alert {
	var message string

	if _, ok := err.(auth.AuthError); ok {
		message = err.Error()
	} else if _, ok := err.(bookshelf.BookshelfError); ok {
		message = err.Error()
	} else {
		log.Println(err)
		message = defaultErrorMessage
	}

	return &Alert{
		Level:   AlertLvError,
		Message: message,
	}
}
