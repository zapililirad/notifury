package user

import "github.com/google/uuid"

type UserFactory struct{}

func (f *UserFactory) CreateUser(firstname, lastname, email string) User {
	// TODO: Add validation
	return User{
		FirstName: firstname,
		LastName:  lastname,
		Email:     email,
		UUID:      uuid.NewString(),
		Active:    false,
	}
}

func (f *UserFactory) CreateUserWithPassword(firstname, lastname, email, pass string) User {
	return User{
		FirstName: firstname,
		LastName:  lastname,
		Email:     email,
		UUID:      uuid.NewString(),
		Active:    true, // TODO: Why?
		password:  SetPassword(pass),
	}
}
