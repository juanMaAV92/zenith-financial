package entities

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uint64    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Code         uuid.UUID `gorm:"column:code;type:uuid;uniqueIndex;not null;default:gen_random_uuid()" json:"code"`
	Username     string    `gorm:"column:username;type:varchar(63);uniqueIndex;not null" json:"username"`
	Email        string    `gorm:"column:email;type:varchar(255);uniqueIndex;not null" json:"email"`
	PasswordHash string    `gorm:"column:password_hash;type:varchar(255);not null" json:"-"`
	PasswordSalt string    `gorm:"column:password_salt;type:varchar(32);not null" json:"-"`
	Currency     string    `gorm:"column:currency;type:varchar(3);not null;default:'USD'" json:"currency"`
	CreatedAt    time.Time `gorm:"column:created_at;type:timestamp with time zone;not null;default:now()" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at;type:timestamp with time zone;not null;default:now()" json:"updated_at"`
}

func (User) TableName() string {
	return "Users"
}
