package users

import "time"

type UserId int32

type User struct {
	Identifier   UserId
	Name         string
	Login        string
	PasswordHash string
	CreatedAt    time.Time
}
