package view

// View all views in this app
type Views struct {
	Books    *View
	EditBook *View
	Error    *View
	Login    *View
	NewBook  *View
	NotFound *View
	Signup   *View
}

// NewViews init all views
func NewViews() *Views {
	return &Views{
		Books:    NewView("books"),
		EditBook: NewView("book"),
		Error:    NewView("error"),
		Login:    NewView("login"),
		NewBook:  NewView("new_book"),
		NotFound: NewView("404"),
		Signup:   NewView("signup"),
	}
}
