package user

type CreateUserDTO struct {
	Name         string `Json:"name"`
	Email        string `Json:"email"`
	PasswordHash string `Json:"password_hash"`
}
