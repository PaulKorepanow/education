package teststore

import (
	"bookLibrary/internal/model"
	"bookLibrary/internal/store"
)

type TestStore struct {
	UserRep UserRep
}

func NewTestStore() *TestStore {
	return &TestStore{}
}

func (s *TestStore) User() store.UserRepository {
	if s.UserRep != nil {
		return &s.UserRep
	}

	s.UserRep = make(map[string]*model.User)

	return &s.UserRep
}

func (s *TestStore) Book() store.BookRepository {
	return nil
}
