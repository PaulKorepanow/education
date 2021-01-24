package model

import (
	"encoding/json"
	"errors"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email             string `gorm:"uniqueIndex"; json:"email"`
	Password          string `gorm:"-"`
	EncryptedPassword string `json:"encrypted_password"`
	Books             []Book `gorm:"foreignKey:UserRefer"`
}

func NewUser(email, password string) *User {
	return &User{
		Email:    email,
		Password: password,
	}
}

func (u *User) BeforeCreation() error {
	if len(u.Password) > 0 {
		enc, err := encryptedString(u.Password)
		if err != nil {
			return err
		}

		u.EncryptedPassword = enc
	}

	return nil
}

func (u *User) CleanPassword() error {
	if u.EncryptedPassword == "" {
		return errors.New("secure password not generated")
	}
	u.Password = ""
	return nil
}

func (u User) Validate() error {
	return validation.ValidateStruct(
		&u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.Required, validation.Length(8, 100)),
	)
}

func (u *User) Marshal() ([]byte, error) {
	return json.Marshal(u)
}

func encryptedString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
