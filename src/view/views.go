package view

import (
	"fmt"
	"io/fs"
	"path/filepath"
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

// NewViews init all views
func NewViews(templates fs.FS) *Views {
	layouts := layoutFiles()

	return &Views{
		Books:    View{name: templateFile("books"), layouts: layouts, templates: templates},
		EditBook: View{name: templateFile("book"), layouts: layouts, templates: templates},
		Error:    View{name: templateFile("error"), layouts: layouts, templates: templates},
		Login:    View{name: templateFile("login"), layouts: layouts, templates: templates},
		NewBook:  View{name: templateFile("new_book"), layouts: layouts, templates: templates},
		NotFound: View{name: templateFile("404"), layouts: layouts, templates: templates},
		Signup:   View{name: templateFile("signup"), layouts: layouts, templates: templates},
	}
}

func layoutFiles() string {
	return filepath.Join("templates", "layouts", "*.gohtml")
}

func templateFile(name string) string {
	filename := fmt.Sprintf("%s.gohtml", name)
	return filepath.Join("templates", filename)
}
