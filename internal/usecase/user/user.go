package user

import (
	"context"
	"errors"
	"go-todo-list/internal/helper"
	"go-todo-list/internal/model"
	"go-todo-list/internal/repository/user"
	"net/http"

	"firebase.google.com/go/v4/auth"
)

type userUsecase struct {
	userRepo user.Repository
	fireAuth *auth.Client
}

func NewUsecase(userRepo user.Repository, fireAuth *auth.Client) UseCase {
	return &userUsecase{
		userRepo: userRepo,
		fireAuth: fireAuth,
	}
}

func (u *userUsecase) SignUp(inputSignUp model.InputUserSignup) (model.User, int, error) {
	userByEmail, err := u.userRepo.GetUserByEmail(inputSignUp.Email)
	if err != nil {
		return userByEmail, http.StatusInternalServerError, err
	}

	if userByEmail.Email != "" {
		return userByEmail, http.StatusConflict, errors.New("email already registered")
	}

	passwordHashed, err := helper.GenerateHashPassword(inputSignUp.Password)
	if err != nil {
		return model.User{}, http.StatusInternalServerError, err
	}

	// save to database
	var newUser model.User
	newUser.Email = inputSignUp.Email
	newUser.Password = passwordHashed

	userSaved, err := u.userRepo.Save(newUser)
	if err != nil {
		return model.User{}, http.StatusInternalServerError, err
	}

	// save to firabase
	paramsNewUser := (&auth.UserToCreate{}).
		Email(inputSignUp.Email).
		Password(passwordHashed)

	_, err = u.fireAuth.CreateUser(context.Background(), paramsNewUser)
	if err != nil {
		return model.User{}, http.StatusInternalServerError, err
	}

	return userSaved, http.StatusOK, nil
}
