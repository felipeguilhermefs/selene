package bookshelf

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

type FullBook struct {
	ID       uint
	UserID   uint
	Title    string
	Author   string
	Comments string
	Tags     string
}

type BookFetcher interface {
	FetchOne(userID uint, id uint) (*FullBook, error)
	FetchMany(userID uint) ([]FullBook, error)
}

type BookRepository interface {
	Insert(book *NewBook) error
	Update(book *UpdatedBook) error
	FindOne(id uint) (*FullBook, error)
	FindMany(userID uint) ([]FullBook, error)
	Delete(id uint) error
}
