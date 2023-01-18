package user

import (
	"context"

	"github.com/google/uuid"
)

type UserService struct {
	storage Storage
}

func (s *UserService) ActivateUser(ctx context.Context, u *User) error {
	u.Active = true
	return nil
}

func (s *UserService) DeactivateUser(ctx context.Context, u *User) error {
	u.Active = false
	return nil
}

func (s *UserService) SetPasswordToUser(ctx context.Context, u *User, p string) error {
	return setPassword(&u.password, p)
}

func (s *UserService) GetUserByUUID(ctx context.Context, uuid string) *User {
	return &User{}
}

func (s *UserService) GetAllUsers(ctx context.Context, limit int, offset int) []*User {
	return []*User{}
}

func (s *UserService) CreateUser(ctx context.Context, firstname, lastname, email string) *User {
	// TODO: Add validation
	return &User{
		FirstName: firstname,
		LastName:  lastname,
		Email:     email,
		UUID:      uuid.NewString(),
		Active:    false,
		password:  password{},
	}
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
