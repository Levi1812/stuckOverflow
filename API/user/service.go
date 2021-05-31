package user

import (
	"errors"
	"fmt"
	"time"

	"github.com/Excellent-Echo/stuckOverflow/API/API/entity"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	SaveNewUser(user entity.UserInput) (UserInputFormat, error)
	LoginUser(input entity.LoginUserInput) (entity.User, error)
}

type userService struct {
	repository UserRepository
}

func UserNewService(repository UserRepository) *userService {
	return &userService{repository}
}

func (s *userService) SaveNewUser(user entity.UserInput) (UserInputFormat, error) {
	genPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)

	if err != nil {
		return UserInputFormat{}, err
	}

	var newUser = entity.User{
		UserName:  user.UserName,
		Email:     user.Email,
		Password:  string(genPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createUser, err := s.repository.CreateUser(newUser)
	formatUser := FormattingUser(createUser)

	if err != nil {
		return formatUser, err
	}

	return formatUser, nil
}

func (s *userService) LoginUser(input entity.LoginUserInput) (entity.User, error) {
	user, err := s.repository.FindByEmail(input.Email)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		newError := fmt.Sprintf("user id %v not found", user.ID)
		return user, errors.New(newError)
	}

	// check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return user, errors.New("invalid password")
	}

	return user, nil
}