package request

type CreateUser struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Currency string `json:"currency"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
