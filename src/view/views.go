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
func NewViews(fileSystem fs.FS) *Views {
	layouts := layoutFiles()

	return &Views{
		Books:    View{name: templateFile("books"), layouts: layouts, fileSystem: fileSystem},
		EditBook: View{name: templateFile("book"), layouts: layouts, fileSystem: fileSystem},
		Error:    View{name: templateFile("error"), layouts: layouts, fileSystem: fileSystem},
		Login:    View{name: templateFile("login"), layouts: layouts, fileSystem: fileSystem},
		NewBook:  View{name: templateFile("new_book"), layouts: layouts, fileSystem: fileSystem},
		NotFound: View{name: templateFile("404"), layouts: layouts, fileSystem: fileSystem},
		Signup:   View{name: templateFile("signup"), layouts: layouts, fileSystem: fileSystem},
	}
}

func layoutFiles() string {
	return filepath.Join("templates", "layouts", "*.gohtml")
}

func templateFile(name string) string {
	filename := fmt.Sprintf("%s.gohtml", name)
	return filepath.Join("templates", filename)
}
