package model

type User struct {
	ID       uint `gorm:"primaryKey"`
	Email    string
	Password string
	Todos    []Todo `gorm:"foreignKey:UserID"`
}
