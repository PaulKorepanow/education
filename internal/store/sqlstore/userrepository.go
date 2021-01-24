package sqlstore

import (
	"bookLibrary/internal/model"
	"fmt"
)

type UserRep struct {
	store *SqlStore
	users map[uint]*model.User
}

func (r *UserRep) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreation(); err != nil {
		return err
	}

	if err := u.CleanPassword(); err != nil {
		return err
	}

	err := r.store.db.Create(u).Error
	if err != nil {
		return err
	}

	r.users[u.ID] = u
	return nil
}

func (r *UserRep) FindByEmail(email string) (*model.User, error) {
	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, fmt.Errorf("user with email(%s) not found", email)
}

func (r *UserRep) FindByID(id uint) (*model.User, error) {
	var user model.User
	if err := r.store.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRep) UpdatePassword(email, password string) (*model.User, error) {
	u, err := r.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	newUser := model.NewUser(email, password)

	if err := newUser.Validate(); err != nil {
		return nil, err
	}

	if err := newUser.BeforeCreation(); err != nil {
		return nil, err
	}

	if err := newUser.CleanPassword(); err != nil {
		return nil, err
	}

	if err := r.store.db.
		Model(&u).
		Update("encrypted_password", newUser.EncryptedPassword).
		Error; err != nil {
		return nil, err
	}

	r.users[newUser.ID] = newUser

	return u, nil
}

func (r *UserRep) DeleteByEmail(email string) error {
	var u model.User
	if err := r.store.db.
		Where("email = ?", email).
		First(&u).
		Delete(&u).Error; err != nil {
		return err
	}
	delete(r.users, u.ID)
	return nil
}

func (r *UserRep) AddBookByEmail(email, title string) (*model.User, error) {
	var u model.User
	if err := r.store.db.
		Where("email = ?", email).
		First(&u). //email unique field
		Association("Books").
		Append(&model.Book{Title: title}); err != nil {
		return nil, err
	}
	r.users[u.ID] = &u
	return &u, nil
}
