package sqlstore

import (
	"gorm.io/gorm"
)

type SqlStore struct {
	db             *gorm.DB
	userRepository *UserRep
}

type Store interface {
	User() UserRepository
}

func NewStore(db *gorm.DB) *SqlStore {
	return &SqlStore{
		db: db,
	}
}

func (s *SqlStore) User() UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRep{store: s}

	return s.userRepository
}
