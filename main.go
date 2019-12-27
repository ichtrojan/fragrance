package main

import (
	"log"
	"net/http"

	"github.com/ichtrojan/fragrance/database"
	"github.com/ichtrojan/fragrance/routes"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}

	database.SeedAdmin()

	if err := http.ListenAndServe(":9000", routes.Init()); err != nil {
		log.Fatal(err)
	}
}
