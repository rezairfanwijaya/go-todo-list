package user

import "go-todo-list/internal/model"

func FormatUserSignUp(user model.User) *model.FormatUserSignUp {
	return &model.FormatUserSignUp{
		ID:    user.ID,
		Email: user.Email,
	}
}
