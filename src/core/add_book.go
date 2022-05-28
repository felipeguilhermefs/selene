package core

type NewBook struct {
	UserID   uint
	Title    string
	Author   string
	Comments string
	Tags     string
}

type BookAdder interface {
	Add(book *NewBook) error
}
