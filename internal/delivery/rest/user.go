package rest

import (
	"go-todo-list/internal/helper"
	"go-todo-list/internal/model"
	"go-todo-list/internal/usecase/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUsecase user.UseCase
}

func NewUserHandler(userUsecase user.UseCase) *UserHandler {
	return &UserHandler{
		userUsecase,
	}
}

func (h *UserHandler) SignUp(c *gin.Context) {
	var inputSignUp model.InputUserSignup

	if err := c.ShouldBindJSON(&inputSignUp); err != nil {
		errBinding := helper.GenerateErrorBinding(err)
		response := helper.GenerateResponseAPI(
			"error",
			http.StatusBadRequest,
			errBinding,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	newUser, httpCode, err := h.userUsecase.SignUp(inputSignUp)
	if err != nil {
		response := helper.GenerateResponseAPI(
			"error",
			httpCode,
			err.Error(),
		)

		c.JSON(httpCode, response)
		return
	}

	userFormatted := user.FormatUserSignUp(newUser)
	response := helper.GenerateResponseAPI(
		"success",
		http.StatusOK,
		userFormatted,
	)

	c.JSON(http.StatusOK, response)
}
