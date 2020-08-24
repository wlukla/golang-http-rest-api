package store

import (
	"golang-http-rest-api/internal/app/model"
)

// UserRepository ...
type UserRepository interface {
	Create(user *model.User) error
	FindByEmail(email string) (*model.User, error)
}