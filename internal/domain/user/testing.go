package user

import "testing"

func TestUser(t *testing.T) *User {
	return &User{
		FirstName: "Ivan",
		LastName:  "Taranov",
		Email:     "user@example.org",
		IsActive:  false,
		UUID:      "f7528611-79cb-4524-a86a-f770e31aa77a",
	}
}
