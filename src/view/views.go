package view

import (
	"fmt"
	"io/fs"
)

// View all views in this app
type Views struct {
	Books    View
	EditBook View
	Error    View
	Login    View
	NewBook  View
	NotFound View
	Signup   View
}

const (
	layoutFiles     = "templates/layouts/*.gohtml"
	templatePattern = "templates/%s.gohtml"
)

// NewViews init all views
func NewViews(templates fs.FS) *Views {
	return &Views{
		Books:    View{name: templateFile("books"), layouts: layoutFiles, templates: templates},
		EditBook: View{name: templateFile("book"), layouts: layoutFiles, templates: templates},
		Error:    View{name: templateFile("error"), layouts: layoutFiles, templates: templates},
		Login:    View{name: templateFile("login"), layouts: layoutFiles, templates: templates},
		NewBook:  View{name: templateFile("new_book"), layouts: layoutFiles, templates: templates},
		NotFound: View{name: templateFile("404"), layouts: layoutFiles, templates: templates},
		Signup:   View{name: templateFile("signup"), layouts: layoutFiles, templates: templates},
	}
}

func templateFile(name string) string {
	return fmt.Sprintf(templatePattern, name)
}
