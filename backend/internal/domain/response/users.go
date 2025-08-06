package response

import (
	"time"

	"github.com/google/uuid"
	"github.com/juanMaAV92/zenith-financial/backend/internal/entities"
)

type User struct {
	Code      uuid.UUID `json:"code"`
	UserName  string    `json:"user_name"`
	Email     string    `json:"email"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at"`
}

type UserLogin struct {
	*User
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func ToUserResponse(user *entities.User) *User {
	return &User{
		Code:      user.Code,
		UserName:  user.Username,
		Email:     user.Email,
		Currency:  user.Currency,
		CreatedAt: user.CreatedAt,
	}
}
