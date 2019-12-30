package database

import (
	"golang.org/x/crypto/bcrypt"
	"log"

	"github.com/ichtrojan/fragrance/models"
)

func SeedAdmin() {
	db := Init()

	db.DropTable(&models.Admin{})

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
