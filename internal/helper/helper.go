package helper

import (
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type ResponseAPI struct {
	Meta meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type meta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func GetEnv(envPath string) (map[string]string, error) {
	env, err := godotenv.Read(envPath)
	if err != nil {
		return env, err
	}

	return env, nil
}

func GenerateErrorBinding(err error) (errorBinding []string) {
	for _, e := range err.(validator.ValidationErrors) {
		errorBinding = append(errorBinding, e.Error())
	}

	return errorBinding
}

func GenerateResponseAPI(message string, code int, data interface{}) *ResponseAPI {
	return &ResponseAPI{
		Meta: meta{
			Code:    code,
			Message: message,
		},
		Data: data,
	}
}

func GenerateHashPassword(rawPassword string) (string, error) {
	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(rawPassword), 10)
	if err != nil {
		return "", err
	}

	return string(passwordHashed), nil
}

func VerifyPassword(rawPassword, hashedPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(rawPassword)); err != nil {
		return err
	}

	return nil
}
