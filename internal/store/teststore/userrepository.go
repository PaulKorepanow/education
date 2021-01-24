package teststore

import (
	"bookLibrary/internal/model"
	"errors"
	"fmt"
	"math/rand"
)

type UserRep map[string]*model.User

func (r UserRep) Create(u *model.User) error {
	if err := u.BeforeCreation(); err != nil {
		return err
	}
	if _, ok := r[u.Email]; ok {
		return errors.New("duplicate email")
	}
	u.ID = uint(rand.Int())
	r[u.Email] = u
	return nil
}

func (r UserRep) FindByEmail(email string) (*model.User, error) {
	if _, ok := r[email]; !ok {
		return nil, errors.New("can not find user by email")
	}
	return r[email], nil
}

func (r UserRep) FindByID(id uint) (*model.User, error) {
	for _, user := range r {
		if user.ID == id {
			return user, nil
		}
	}
	return nil, fmt.Errorf("user with id(%d) not found", id)
}

func (r UserRep) UpdatePassword(email, password string) (*model.User, error) {
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

func (r UserRep) AddBookByEmail(email, title string) (*model.User, error) {
	return nil, nil
}
