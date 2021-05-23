package view

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
func NewViews() *Views {
	return &Views{
		Books:    View{name: "books"},
		EditBook: View{name: "book"},
		Error:    View{name: "error"},
		Login:    View{name: "login"},
		NewBook:  View{name: "new_book"},
		NotFound: View{name: "404"},
		Signup:   View{name: "signup"},
	}
}
