package request

type CreateUser struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Currency string `json:"currency"`
}

type ValidateUserCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
