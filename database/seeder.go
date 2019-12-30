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

func SeedCategory() {
	db := Init()

	db.DropTable(&models.Category{})

	db.AutoMigrate(&models.Category{})

	categories := []models.Category{
		{
			Name:  "Feminine",
			Slug:  "feminine",
			Image: "https://res.cloudinary.com/ichtrojan/image/upload/v1577708397/fragrance/feminine-bg.jpg",
		},
		{
			Name:  "Masculine",
			Slug:  "masculine",
			Image: "https://res.cloudinary.com/ichtrojan/image/upload/v1577708255/fragrance/masculine-bg.jpg",
		},
		{
			Name:  "Unisex",
			Slug:  "unisex",
			Image: "https://res.cloudinary.com/ichtrojan/image/upload/v1577708407/fragrance/unisex-bg.jpg",
		},
	}

	for _, category := range categories {
		db.Create(&category)
	}
}

func hashAndSalt(password []byte) string {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)

	if err != nil {
		log.Println(err)
	}

	return string(hash)
}
