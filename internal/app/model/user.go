package model

import (
	"golang.org/x/crypto/bcrypt"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// User ...
type User struct {
	ID                string
	Email             string
	Password string
	EncryptedPassword string
}

func encryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)

	if err != nil {
		return "", err
	}

	return string(b), nil;
}

// BeforeCreate ...
func (u *User) BeforeCreate() error {
	if (len(u.Password) > 0) {
		enc, err := encryptString(u.Password);

		if err != nil {
			return err;
		}

		u.EncryptedPassword = enc
	}

	return nil
}

// Validate ...
func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.Required, validation.Length(6, 100)),
	)
}