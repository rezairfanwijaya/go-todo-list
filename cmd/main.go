package main

import (
	"context"
	"go-todo-list/internal/database"
	"go-todo-list/internal/delivery/rest"
	"log"

	firebase "firebase.google.com/go/v4"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

func main() {
	// db connection
	db, err := database.NewConnection("../.env")
	if err != nil {
		log.Fatal(err)
	}

	// firebase
	opt := option.WithCredentialsFile("../firebase_key.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatal(err)
	}

	// Access Auth service from default app
	defaultClient, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	// server
	router := gin.Default()

	rest.InitRouter(db, router, defaultClient)

	if err := router.Run(":9797"); err != nil {
		log.Fatalf("failed start server: %v\n", err)
	}
}
