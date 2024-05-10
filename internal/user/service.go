package user

import "context"

type Service interface {
	CreateUser() //todo: понять что делает интерфейс

}

type service struct {
	storage Storage
}

func (s *service) CreateUser(ctx context.Context, name, email, password_hash string) {

}

func (s *service)