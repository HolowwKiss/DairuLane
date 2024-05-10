package user

// usage with database
type Storage interface {
	CreateUser(name, email, password_hash string) *ContentCreator
}
