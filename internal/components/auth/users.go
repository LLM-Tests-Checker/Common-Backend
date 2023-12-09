package auth

import "fmt"

type UserProvider interface {
	ProvideUserById(userId int32) (error, *User)

	ProvideUserByLogin(login string) (error, *User)
}

type User struct {
	Identifier   int32
	Name         string
	Login        string
	PasswordHash string
}

func NewUserProvider() UserProvider {
	return defaultUserProvider{}
}

type defaultUserProvider struct {
}

func (provider defaultUserProvider) ProvideUserById(userId int32) (error, *User) {
	return fmt.Errorf("not implemented yet"), nil
}

func (provider defaultUserProvider) ProvideUserByLogin(login string) (error, *User) {
	return fmt.Errorf("not implemented yet"), nil
}
