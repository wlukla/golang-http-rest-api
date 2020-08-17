package store_test

import (
	"golang-http-rest-api/internal/app/model"
	"golang-http-rest-api/internal/app/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	u, err := s.User().Create(&model.User{
		Email: "user@example.com",
	})

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	email := "user@example.com"

	_, err := s.User().FindByEmail(email)

	assert.Error(t, err)

	s.User().Create(&model.User{
		Email: "user@example.com",
	})

	u, err := s.User().FindByEmail(email)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}
