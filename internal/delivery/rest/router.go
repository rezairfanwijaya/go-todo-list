package rest

import (
	userrepo "go-todo-list/internal/repository/user"
	userusecase "go-todo-list/internal/usecase/user"

	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(dbConnection *gorm.DB, router *gin.Engine, authClient *auth.Client) {
	userRepository := userrepo.NewRepository(dbConnection)
	userUsecase := userusecase.NewUsecase(userRepository, authClient)
	userDelivery := NewUserHandler(userUsecase)

	router.POST("/signup", userDelivery.SignUp)
	router.POST("/signin", userDelivery.SignIn)
}
