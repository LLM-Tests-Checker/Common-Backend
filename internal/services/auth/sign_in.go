package auth

type SignInService interface {
	PerformSignIn(login, passwordHash string) (error, User)
}
