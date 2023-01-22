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
			p := password{}
			setPassword(&p, tc.password)
			assert.NotEqual(t, tc.password, p.password)
			assert.NotEmpty(t, p.password)
			assert.IsType(t, time.Now(), p.created)
			assert.NotEqual(t, &time.Time{}, p.created)

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

	p := password{}
	setPassword(&p, "P@$$w0rd")

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
