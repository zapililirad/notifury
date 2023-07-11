package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/zapililirad/notifury/internal/app"
)

type UserService struct {
	storage Storage
}

func (s *UserService) ActivateUser(ctx context.Context, u *User) error {
	u.IsActive = true
	return nil
}

func (s *UserService) DeactivateUser(ctx context.Context, u *User) error {
	u.IsActive = false
	return nil
}

func (s *UserService) SetPasswordToUser(ctx context.Context, u *User, p string) error {
	return setPassword(&u.password, p)
}

func (s *UserService) GetUserByUUID(ctx context.Context, uuid string) (*User, error) {
	return nil, app.ErrNotImplemented
}

func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	return nil, app.ErrNotImplemented
}

func (s *UserService) GetAllUsers(ctx context.Context, limit int, offset int) []*User {
	return nil
}

func (s *UserService) CreateUser(ctx context.Context, firstname, lastname, email string) (*User, error) {
	// access.Needs.CreateUser
	if err := validateUser(firstname, lastname, email); err != nil {
		return nil, err
	}

	return &User{
		FirstName: firstname,
		LastName:  lastname,
		Email:     email,
		UUID:      uuid.NewString(),
		IsActive:  false,
	}, nil
}

// type Service interface {
// 	GetUserByUUID(ctx context.Context, uuid string) *User
// 	GetAllUsers(ctx context.Context, limit int, offset int) []*User
// }

// type service struct {
// 	storage Storage
// }

// func (s *service) GetUserByUUID(ctx context.Context, uuid string) *User {
// 	return &User{}
// }

// func (s *service) GetAllUsers(ctx context.Context, limit int, offset int) []*User {
// 	return []*User{}
// }
