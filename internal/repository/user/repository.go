package user

import "go-todo-list/internal/model"

type Repository interface {
	Save(newUser model.User) (model.User, error)
	GetUserByEmail(email string) (model.User, error)
	GetUserByID(id uint) (model.User, error)
}
