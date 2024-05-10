package user

import "context"

type Service interface {
	CreateUser(ctx context.Context, dto *CreateUserDTO) *User
}

type service struct {
	storage Storage
}

func (s *service) CreateUser(ctx context.Context, dto *CreateUserDTO) *User {
	return nil
}
