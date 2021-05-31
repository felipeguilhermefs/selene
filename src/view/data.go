package view

// Data to enrich templates
type Data struct {
	Dynamic DynamicData
	Static  StaticData
}

// DynamicData that varies according to template and render
type DynamicData struct {
	Alert   *Alert
	User    interface{}
	Content interface{}
}

// StaticData that remains between templates and renders
type StaticData struct {
	Scripts []Dependency
	Styles  []Dependency
}

// WithError sets errors that will be shown to user
func (d *Data) WithError(err error) *Data {
	d.Dynamic.Alert = newErrorAlert(err)
	return d
}

// NewData creates a new data with content to be yield to templates
func NewData(content interface{}) *Data {
	return &Data{
		Dynamic: DynamicData{
			Content: content,
		},
	}
}
