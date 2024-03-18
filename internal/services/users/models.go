package users

import "time"

type UserId int32

type User struct {
	Identifier   UserId
	Login        string
	PasswordHash string
	CreatedAt    time.Time
}

func (userId UserId) Int32() int32 {
	return int32(userId)
}

func (userId UserId) Int() int {
	return int(userId)
}
