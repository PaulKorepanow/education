package store

import "bookLibrary/internal/model"

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	if err := u.BeforeCreation(); err != nil {
		return nil, err
	}

	if err := u.Validate(); err != nil {
		return nil, err
	}

	err := r.store.db.Create(u).Error
	if err != nil {
		return nil, err
	}
	var user model.User
	if err := r.store.db.Last(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	var u model.User
	err := r.store.db.Where("email = ?", email).First(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) UpdatePasswordByEmail(email, password string) (*model.User, error) {
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

func (r *UserRepository) DeleteByEmail(email string) error {
	var u model.User
	if err := r.store.db.Where("email = ?", email).Delete(&u).Error; err != nil {
		return err
	}
	return nil
}
