package database

import (
	"github.com/ichtrojan/fragrance/models"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func SeedAdmin() {
	db := Init()

	db.AutoMigrate(&models.Admin{})

	admin := models.Admin{
		Name:     "Cleopatra",
		Email:    "cleopatra@gofragrance.xyz",
		Password: hashAndSalt([]byte("iloveegypt")),
	}

	db.Create(&admin)
}

func hashAndSalt(password []byte) string {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)

	if err != nil {
		log.Println(err)
	}

	return string(hash)
}
