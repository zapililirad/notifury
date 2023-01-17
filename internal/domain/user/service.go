package user

import "context"

type Service interface {
	GetUserByUUID(ctx context.Context, uuid string) *User
	GetAllUsers(ctx context.Context, limit int, offset int) []*User
}

type service struct {
	storage Storage
}

func (s *service) GetUserByUUID(ctx context.Context, uuid string) *User {
	return &User{}
}

func (s *service) GetAllUsers(ctx context.Context, limit int, offset int) []*User {
	return []*User{}
}
