package auth

type UserProvider interface {
	ProvideUserById(userId int32) (error, User)

	ProvideUserByLogin(login string) (error, User)
}

type User struct {
	Identifier   int32
	Name         string
	Login        string
	PasswordHash string
}
