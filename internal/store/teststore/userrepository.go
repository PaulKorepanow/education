package teststore

import (
	"bookLibrary/internal/model"
	"errors"
)

type UserRep map[string]*model.User

func (r UserRep) Create(u *model.User) error {
	if _, ok := r[u.Email]; ok {
		return errors.New("duplicate email")
	}
	r[u.Email] = u
	return nil
}

func (r UserRep) FindByEmail(email string) (*model.User, error) {
	if _, ok := r[email]; !ok {
		return nil, errors.New("can not find user by email")
	}
	return r[email], nil
}

func (r UserRep) UpdatePasswordByEmail(email, password string) (*model.User, error) {
	if _, ok := r[email]; !ok {
		return nil, errors.New("can not find user by email")
	}
	u := r[email]
	u.Password = password
	return u, nil
}

func (r UserRep) DeleteByEmail(email string) error {
	if _, ok := r[email]; !ok {
		return errors.New("can not find user by email")
	}
	delete(r, email)
	return nil
}
