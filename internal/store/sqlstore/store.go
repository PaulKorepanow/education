package sqlstore

import (
	"bookLibrary/internal/model"
	"bookLibrary/internal/store"
	"gorm.io/gorm"
)

type SqlStore struct {
	db             *gorm.DB
	userRepository *UserRep
	bookRepository *BookRepository
}

func NewStore(db *gorm.DB) *SqlStore {
	if err := db.AutoMigrate(&model.User{}, &model.Book{}); err != nil {
		panic(err)
	}
	return &SqlStore{
		db: db,
	}
}

func (s *SqlStore) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRep{
		store: s,
	}

	return s.userRepository
}

func (s *SqlStore) Book() store.BookRepository {
	if s.bookRepository != nil {
		return s.bookRepository
	}

	s.bookRepository = &BookRepository{
		store: s,
	}

	return s.bookRepository
}
