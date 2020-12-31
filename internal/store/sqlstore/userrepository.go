package sqlstore

import (
	"bookLibrary/internal/model"
)

type UserRep struct {
	store *SqlStore
}

func (r *UserRep) Create(u *model.User) error {
	if err := u.BeforeCreation(); err != nil {
		return err
	}

	if err := u.Validate(); err != nil {
		return err
	}

	err := r.store.db.Create(u).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRep) FindByEmail(email string) (*model.User, error) {
	var u model.User
	err := r.store.db.Where("email = ?", email).First(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRep) UpdatePasswordByEmail(email, password string) (*model.User, error) {
	u, err := r.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	//FIXME: only password field
	u.Email = email
	u.Password = password

	if err := u.BeforeCreation(); err != nil {
		return nil, err
	}

	if err := r.store.db.Save(u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

func (r *UserRep) DeleteByEmail(email string) error {
	var u model.User
	if err := r.store.db.Where("email = ?", email).Delete(&u).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRep) AddBookByEmail(email, title string) (*model.User, error) {
	var u model.User
	if err := r.store.db.Model(&u).
		Where("email = ?", email).
		Association("Books").
		Append(&model.Book{Title: title}); err != nil {
		return nil, err
	}
	return &u, nil
}
