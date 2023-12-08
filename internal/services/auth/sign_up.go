package auth

type SignUpService interface {
	PerformSignUp(name, login, passwordHash string) (error, User)
}
