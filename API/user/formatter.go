package user

import (
	"time"

	"github.com/Excellent-Echo/stuckOverflow/API/API/entity"
)

type UserInputFormat struct {
	UserID   uint32 `json:"id"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
}

// type UserFormat struct {
// 	UserID   uint32            `json:"id"`
// 	UserName string            `json:"user_name"`
// 	Email    string            `json:"email"`
// 	User     entity.UserDetail `json:"user_detail"`
// }

type UserFormat struct {
	ID        uint32 `json:"id"`
	UserName  string `json:"user_name"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`
	Location  string `json:"location"`
}

type DeleteFormat struct {
	Message    string    `json:"message"`
	TimeDelete time.Time `json:"delete_time"`
}

func FormattingUserInput(user entity.User) UserInputFormat {
	userFormat := UserInputFormat{
		UserID:   user.ID,
		UserName: user.UserName,
		Email:    user.Email,
	}

	return userFormat
}

func FormattingUser(user entity.User) UserFormat {
	userFormat := UserFormat{
		ID:        user.ID,
		UserName:  user.UserName,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Avatar:    user.Avatar,
		Location:  user.Location,
	}

	return userFormat
}

func FormatDelete(msg string) DeleteFormat {
	var deleteFormat = DeleteFormat{
		Message:    msg,
		TimeDelete: time.Now(),
	}
	return deleteFormat
}
