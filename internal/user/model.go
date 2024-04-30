package user

type ContentCreator struct {
	UUID          int    `Json:"uuid"`
	Name          string `Json:"name"`
	Email         string `Json:"email"`
	PasswordHash  string `Json:"password_hash"`
	Subscribers   int    `Json:"subscribers"`
	Subscriptions string `Json:"subscriptions"`
	Articles      string `Json:"articles"`
	AutorRating   int    `Json:"autor_rating"`
}
