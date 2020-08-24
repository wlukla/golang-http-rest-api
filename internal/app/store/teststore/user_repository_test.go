package teststore_test

import (
	"golang-http-rest-api/internal/app/model"
	"golang-http-rest-api/internal/app/store/teststore"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s := teststore.New()
	u := model.TestUser(t);

	err := s.User().Create(u)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s := teststore.New()

	email := "user@example.com"

	_, err := s.User().FindByEmail(email)

	assert.Error(t, err)

	u := model.TestUser(t)
	u.Email = email;

	s.User().Create(u)

	u, err = s.User().FindByEmail(email)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}
