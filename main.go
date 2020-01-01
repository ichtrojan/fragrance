package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ichtrojan/fragrance/database"
	"github.com/ichtrojan/fragrance/routes"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}

	database.SeedAdmin()

	database.SeedCategory()

	database.SeedScent()

	database.SeedBottleSizes()

	database.SeedBottle()

	port, exist := os.LookupEnv("PORT")

	if !exist {
		log.Fatal("PORT not set in .env")
	}

	if err := http.ListenAndServe(":"+port, routes.Init()); err != nil {
		log.Fatal(err)
	}
}
