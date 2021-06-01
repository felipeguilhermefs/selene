package view

// Data to enrich templates
type Data struct {
	Alert   *Alert
	User    interface{}
	Content interface{}
}

// WithError sets errors that will be shown to user
func (d *Data) WithError(err error) *Data {
	d.Alert = newErrorAlert(err)
	return d
}

// NewData creates a new data with content to be yield to templates
func NewData(content interface{}) *Data {
	return &Data{
		Content: content,
	}
}
