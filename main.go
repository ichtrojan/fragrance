package main

import (
	"errors"
	"github.com/ichtrojan/thoth"
	"log"
	"net/http"
	"os"

	"github.com/ichtrojan/fragrance/database"
	"github.com/ichtrojan/fragrance/routes"
	"github.com/joho/godotenv"
)

func main() {
	logger, _ := thoth.Init("log")

	if err := godotenv.Load(); err != nil {
		logger.Log(errors.New("no .env file found"))
		log.Fatal("No .env file found")
	}

	database.SeedAdmin()

	database.SeedCategory()

	database.SeedScent()

	database.SeedBottleSizes()

	database.SeedBottle()

	port, exist := os.LookupEnv("PORT")

	if !exist {
		logger.Log(errors.New("PORT not set in .env"))
		log.Fatal("PORT not set in .env")
	}

	if err := http.ListenAndServe(":"+port, routes.Init()); err != nil {
		logger.Log(err)
		log.Fatal(err)
	}
}
