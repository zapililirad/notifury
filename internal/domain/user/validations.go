package user

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

func validateUser(firstName, lastName, email string) error {
	user := struct {
		firstName string
		lastName  string
		email     string
	}{
		firstName: firstName,
		lastName:  lastName,
		email:     email,
	}

	return validation.ValidateStruct(
		&user,
		validation.Field(&user.firstName, validation.Required, is.PrintableASCII),
		validation.Field(&user.lastName, validation.Required, is.PrintableASCII),
		validation.Field(&user.email, validation.Required, is.Email),
	)
}
