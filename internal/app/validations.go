package app

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

func ValidateName(name string) error {
	return validation.Validate(
		name,
		validation.Required,
		is.PrintableASCII,
	)
}
