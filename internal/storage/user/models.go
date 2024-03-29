package user

const (
	userFieldLogin      = "login"
	userFieldIdentifier = "identifier"
)

type user struct {
	Identifier   int32  `bson:"identifier"`
	Login        string `bson:"login"`
	PasswordHash string `bson:"password_hash"`
	CreatedAt    string `bson:"created_at"`
}
