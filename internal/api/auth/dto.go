package auth

type SignUpRequest struct {
	UserName         string `json:"user_name"`
	UserLogin        string `json:"user_login"`
	UserPasswordHash string `json:"user_password_hash"`
}

type SignInRequest struct {
	UserLogin        string `json:"user_login"`
	UserPasswordHash string `json:"user_password_hash"`
}
