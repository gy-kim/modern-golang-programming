package main

import (
	"log"

	"github.com/gy-kim/modern-golang-programming/dino/databaselayer"
	"github.com/gy-kim/modern-golang-programming/dino/dinowebportal/dinoapi"
)

func main() {
	db, err := databaselayer.GetDatabaseHandler(databaselayer.MONGODB, "mongodb://127.0.0.1")
	if err != nil {
		log.Fatal(err)
	}
	dinoapi.RunApi("localhost:8080", db)
}
