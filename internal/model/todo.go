package model

type StatusTodo string

type Todo struct {
	ID          uint `gorm:"primaryKey"`
	UserID      string
	Title       string
	Description string
	Status      StatusTodo
}
