package user

type Storage interface {
	CreateUser(name, email, password_hash string) *ContentCreator
}
