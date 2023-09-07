package main

import (
	"go-todo-list/internal/database"
	"log"
)

func main() {
	// db connection
	db, err := database.NewConnection("../.env")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(db)
}
