package models

import (
	"gorm.io/gorm"

	"github.com/gbalves1989/projeto-api-golang/http/responses"
)

type User struct {
	gorm.Model
	Name         string `json:"name"`
	Email        string `gorm:"uniqueIndex" json:"email"`
	PasswordHash string `json:"-"`
	AvatarURL    string `json:"avatar_url"`
}

func (u *User) ToUserResponse() responses.UserData {
	return responses.UserData{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		AvatarURL: u.AvatarURL,
	}
}
