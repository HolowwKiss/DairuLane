package user

import "context"

type Service interface {
	CreateUser(ctx context.Context)
}

type service struct {
	storage Storage
}

func (s *service) CreateUser(ctx context.Context, name, email, password_hash string) {

}
