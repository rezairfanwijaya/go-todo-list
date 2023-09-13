package user

import (
	"go-todo-list/internal/model"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) Save(newUser model.User) (model.User, error) {
	if err := ur.db.Create(&newUser).Error; err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (ur *userRepository) GetUserByEmail(email string) (model.User, error) {
	var user model.User

	if err := ur.db.Where("email = ?", email).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (ur *userRepository) GetUserByID(id uint) (model.User, error) {
	var user model.User

	if err := ur.db.Where("id = ?", id).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
