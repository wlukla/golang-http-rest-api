package teststore

import (
	"golang-http-rest-api/internal/app/model"
	"golang-http-rest-api/internal/app/store"
)

// UserRepository ...
type UserRepository struct {
	store *Store
	users map[string]*model.User
}

// Create ...
func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	u.ID = string(len(r.users) + 1)
	r.users[u.ID] = u

	return nil
}

// FindByEmail ...
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u, ok := r.users[email]

	if !ok {
		return nil, store.ErrRecordNotFound
	}

	return u, nil
}

// Find ...
func (r *UserRepository) Find(id string) (*model.User, error) {
	u, ok := r.users[id]

	if !ok {
		return nil, store.ErrRecordNotFound
	}

	return u, nil
}
