package store

import "bookLibrary/internal/model"

type UserRepository interface {
	Create(u *model.User) error
	FindByEmail(email string) (*model.User, error)
	UpdatePasswordByEmail(email, password string) (*model.User, error)
	DeleteByEmail(email string) error
}

type BookRepository interface {
	Create(b *model.Book) error
	FindByTitle(title string) (*model.Book, error)
}
