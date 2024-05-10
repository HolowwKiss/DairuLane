package user

// usage with database
type Storage interface {
	CreateUser(user *User) *User
}
