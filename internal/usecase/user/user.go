package user

import (
	"context"
	"errors"
	"go-todo-list/internal/helper"
	"go-todo-list/internal/model"
	"go-todo-list/internal/repository/user"
	"net/http"
	"strconv"

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

func (u *userUsecase) SignIn(inputSignIn model.InputSignIn) (model.User, string, int, error) {
	var customToken string
	userByEmail, err := u.userRepo.GetUserByEmail(inputSignIn.Email)
	if err != nil {
		return model.User{}, customToken, http.StatusInternalServerError, err
	}

	if userByEmail.ID == 0 {
		return model.User{}, customToken, http.StatusNotFound, errors.New("email not registered")
	}

	if err := helper.VerifyPassword(inputSignIn.Password, userByEmail.Password); err != nil {
		return model.User{}, customToken, http.StatusBadRequest, errors.New("invalid password")
	}

	stringID := strconv.Itoa(int(userByEmail.ID))
	customToken, err = u.fireAuth.CustomToken(context.Background(), stringID)
	if err != nil {
		return model.User{}, customToken, http.StatusInternalServerError, err
	}

	return userByEmail, customToken, http.StatusOK, nil
}
