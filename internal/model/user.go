package model

type User struct {
	ID       uint `gorm:"primaryKey"`
	Email    string
	Password string
	Todos    []Todo `gorm:"foreignKey:UserID"`
}

type InputUserSignup struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type FormatUserSignUp struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
}
