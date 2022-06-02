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

type BookRemover interface {
	Remove(userID uint, id uint) error
}
