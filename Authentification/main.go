package main

import (
	"Authentification/router"
	"github.com/joho/godotenv"
	"log"
)

func main() {

	env := godotenv.Load(".env")
	if env != nil {
		log.Fatalf("Some error occured. Err: %s", env)
	}

	router.Router()
}
