package user

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUserService_ActivateUser(t *testing.T) {
	s := &UserService{}

	u := TestUser(t)
	ctx := context.Background()

	assert.NoError(t, s.ActivateUser(ctx, u))
	assert.Equal(t, true, u.IsActive)
}

func TestUserService_DeactivateUser(t *testing.T) {
	s := &UserService{}

	u := TestUser(t)
	ctx := context.Background()

	u.IsActive = true

	assert.NoError(t, s.DeactivateUser(ctx, u))
	assert.Equal(t, false, u.IsActive)
}

func TestUserService_SetPasswordToUser(t *testing.T) {
	s := &UserService{}

	testCases := []struct {
		name     string
		password string
		ok       bool
	}{
		{
			name:     "Valid password",
			password: "P@$$w0rd",
			ok:       true,
		},
		{
			name:     "Empty password",
			password: "",
			ok:       false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			u := TestUser(t)
			ctx := context.Background()

			assert.NoError(t, s.SetPasswordToUser(ctx, u, tc.password))
			assert.NotEqual(t, tc.password, u.password.password)
			assert.NotEmpty(t, u.password.password)
			assert.NotEqual(t, &time.Time{}, u.password.created)

		})
	}
}

func TestUserService_CreateUser(t *testing.T) {
	s := &UserService{}

	testCases := []struct {
		name      string
		firstname string
		lastname  string
		email     string
		ok        bool
	}{
		{
			name:      "Valid user",
			firstname: "Ivan",
			lastname:  "Taranov",
			email:     "user@example.org",
			ok:        true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			u, err := s.CreateUser(ctx, tc.firstname, tc.lastname, tc.email)

			assert.NoError(t, err)
			assert.Equal(t, tc.firstname, u.FirstName)
			assert.Equal(t, tc.lastname, u.LastName)
			assert.Equal(t, tc.email, u.Email)
			assert.Equal(t, false, u.IsActive)

			assert.NotPanics(t, func() { uuid.MustParse(u.UUID) }, "User id is not UUID")
		})
	}

}
