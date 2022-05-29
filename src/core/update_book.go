package core

type UpdatedBook struct {
	ID       uint
	UserID   uint
	Title    string
	Author   string
	Comments string
	Tags     string
}

type BookUpdater interface {
	Update(book *UpdatedBook) error
}
