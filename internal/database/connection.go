package database

import (
	"fmt"
	"go-todo-list/internal/helper"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewConnection(envPath string) (*gorm.DB, error) {
	env, err := helper.GetEnv(envPath)
	if err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		env["DB_USERNAME"],
		env["DB_PASSWORD"],
		env["DB_HOST"],
		env["DB_PORT"],
		env["DB_NAME"],
	)
	
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}

	return db, nil
}
