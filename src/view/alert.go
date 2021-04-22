package view

const (
	// AlertLvError error alert level
	AlertLvError = "danger"
	// AlertLvWarning  warn alert level
	AlertLvWarning = "warning"
	// AlertLvInfo  info alert level
	AlertLvInfo = "info"
	// AlertLvSuccess  success alert level
	AlertLvSuccess = "success"
)

// Alert is used to render alert messages in templates
type Alert struct {
	Level   string
	Message string
}

func newErrorAlert(err error) *Alert {
	return &Alert{
		Level:   AlertLvError,
		Message: err.Error(),
	}
}
