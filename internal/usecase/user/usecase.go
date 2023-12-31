package user

import "go-todo-list/internal/model"

type UseCase interface {
	SignUp(inputSignUp model.InputUserSignup) (model.User, int, error)
	SignIn(inputSignIn model.InputSignIn) (model.User, string, int, error)
}
