package user

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/zapililirad/notifury/internal/domain/access"
)

func TestUser_GetName(t *testing.T) {
	u := User{
		FirstName: "Ivan",
		LastName:  "Petrov",
	}
	assert.Equal(t, "Ivan Petrov", u.GetName())
}

func TestUser_GetSecurityUUID(t *testing.T) {
	u := User{
		UUID: uuid.NewString(),
	}
	assert.Equal(t, u.UUID, u.GetSecurityUUID())
}

func TestUser_GetSecurityPrincipalType(t *testing.T) {
	u := User{}
	assert.Equal(t, access.User, u.GetSecurityPrincipalType())
}

func TestUser_SetPassword(t *testing.T) {
	testCases := []struct {
		name     string
		password string
	}{
		{
			name:     "Valid password",
			password: "P@$$w0rd",
		},
		{
			name:     "Empty password",
			password: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p := SetPassword(tc.password)
			assert.NotEqual(t, tc.password, p.password)
			assert.NotEmpty(t, p.password)
			assert.IsType(t, time.Now(), p.created)
		})
	}
}

func TestUser_PasswordCompareWithString(t *testing.T) {
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
		{
			name:     "Invalid password",
			password: "Invalid",
			ok:       false,
		},
	}

	p := SetPassword("P@$$w0rd")

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.ok {
				assert.NoError(t, p.CompareWithString(tc.password))
			} else {
				assert.Error(t, p.CompareWithString(tc.password))
			}
		})
	}
}

func TestUser_CreateUser(t *testing.T) {
	f := UserFactory{}
	u := f.CreateUser(
		"Ivan",
		"Petrov",
		"ivan@example.org",
	)

	assert.Equal(t, "Ivan", u.FirstName)
	assert.Equal(t, "Petrov", u.LastName)
	assert.Equal(t, "ivan@example.org", u.Email)

	_, err := uuid.Parse(u.UUID)
	assert.NoError(t, err)
}

func TestUser_CreateUserWithPassword(t *testing.T) {
	f := UserFactory{}
	u := f.CreateUserWithPassword(
		"Ivan",
		"Petrov",
		"ivan@example.org",
		"P@$$w0rd",
	)

	assert.Equal(t, "Ivan", u.FirstName)
	assert.Equal(t, "Petrov", u.LastName)
	assert.Equal(t, "ivan@example.org", u.Email)

	_, err := uuid.Parse(u.UUID)
	assert.NoError(t, err)

	assert.IsType(t, password{}, u.password)

	assert.NoError(t, u.password.CompareWithString("P@$$w0rd"))
}
