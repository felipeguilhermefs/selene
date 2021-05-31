package view

var (
	jquery = Dependency{
		URL:      "https://code.jquery.com/jquery-3.5.1.slim.min.js",
		Checksum: "sha384-DfXdz2htPH0lsSSs5nCTpuj/zy4C+OGpamoFVy38MVBnE+IbbVYUew+OrCXaRkfj",
	}

	bootstrapJS = Dependency{
		URL:      "https://cdn.jsdelivr.net/npm/bootstrap@4.6.0/dist/js/bootstrap.bundle.min.js",
		Checksum: "sha384-Piv4xVNRyMGpqkS2by6br4gNJ7DXjqk09RmUpJ8jgGtD7zP9yug3goQfGII0yAns",
	}

	bootstrapCSS = Dependency{
		URL:      "https://cdn.jsdelivr.net/npm/bootstrap@4.6.0/dist/css/bootstrap.min.css",
		Checksum: "sha384-B0vP5xmATw1+K9KRQjQERJvTumQW0nPEzvF6L/Z6nronJ3oUOFUFpCjEUQouq2+l",
	}

	allowedScripts = []Dependency{jquery, bootstrapJS}
	allowedStyles  = []Dependency{bootstrapCSS}
)

// Dependency holds data that identifies external dependencies
type Dependency struct {
	URL      string
	Checksum string
}
