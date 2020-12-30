package sqlstore

import (
	"bookLibrary/internal/store"
	"gorm.io/gorm"
)

type SqlStore struct {
	db             *gorm.DB
	userRepository *UserRep
}

func NewStore(db *gorm.DB) *SqlStore {
	return &SqlStore{
		db: db,
	}
}

func (s *SqlStore) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRep{store: s}

	return s.userRepository
}
