package view

var (
	jquery = Script{
		URL:      "https://code.jquery.com/jquery-3.5.1.slim.min.js",
		Checksum: "sha384-DfXdz2htPH0lsSSs5nCTpuj/zy4C+OGpamoFVy38MVBnE+IbbVYUew+OrCXaRkfj",
	}
	bootstrap = Script{
		URL:      "https://cdn.jsdelivr.net/npm/bootstrap@4.6.0/dist/js/bootstrap.bundle.min.js",
		Checksum: "sha384-Piv4xVNRyMGpqkS2by6br4gNJ7DXjqk09RmUpJ8jgGtD7zP9yug3goQfGII0yAns",
	}

	allowedScripts = []Script{jquery, bootstrap}
)

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
	Scripts []Script
}

// Script holds data that identifies script dependencies
type Script struct {
	URL      string
	Checksum string
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
