package teststore

import (
	"bookLibrary/internal/model"
	"bookLibrary/internal/store/sqlstore"
)

type TestStore struct {
	UserRep UserRep
}

func NewTestStore() *TestStore {
	return &TestStore{}
}

func (s *TestStore) User() sqlstore.UserRepository {
	if s.UserRep != nil {
		return &s.UserRep
	}

	s.UserRep = make(map[string]*model.User)

	return &s.UserRep
}
