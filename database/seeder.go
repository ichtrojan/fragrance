package database

import (
	"log"

	"golang.org/x/crypto/bcrypt"

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

func SeedBottleSizes() {
	db := Init()

	db.DropTable(&models.BottleSize{})

	db.AutoMigrate(&models.BottleSize{})

}

func SeedScent() {
	db := Init()

	db.DropTable(&models.Scent{})

	db.AutoMigrate(&models.Scent{})

	scents := []models.Scent{
		{
			Name:  "Woodsy",
			Slug:  "woodsy",
			Image: "https://res.cloudinary.com/ichtrojan/image/upload/v1577750699/fragrance/slide1.png",
			Price: 2.0,
		},
		{
			Name:  "Fruity",
			Slug:  "fruity",
			Image: "https://res.cloudinary.com/ichtrojan/image/upload/v1577750697/fragrance/slide5.png",
			Price: 4.0,
		},
		{
			Name:  "Floral",
			Slug:  "floral",
			Image: "https://res.cloudinary.com/ichtrojan/image/upload/v1577750698/fragrance/slide2.png",
			Price: 4.0,
		},
		{
			Name:  "Earthly",
			Slug:  "earthly",
			Image: "https://res.cloudinary.com/ichtrojan/image/upload/v1577750698/fragrance/slide3.png",
			Price: 2.5,
		},
		{
			Name:  "Oceanic",
			Slug:  "oceanic",
			Image: "https://res.cloudinary.com/ichtrojan/image/upload/v1577750699/fragrance/slide4.png",
			Price: 3.5,
		},
	}

	for _, scent := range scents {
		db.Create(&scent)
	}
}

func SeedBottle() {
	db := Init()

	db.DropTable(&models.Bottle{})

	db.AutoMigrate(&models.Bottle{})

	bottles := []models.Bottle{
		{
			Name:  "Bullet",
			Slug:  "bullet",
			Image: "https://res.cloudinary.com/ichtrojan/image/upload/v1577751580/fragrance/perfume4.png",
			BottleSizes: []models.BottleSize{
				{
					Size:  "30 ml",
					Price: 4.0,
				},
				{
					Size:  "60 ml",
					Price: 8.0,
				},
				{
					Size:  "90 ml",
					Price: 12.5,
				},
			},
		},
		{
			Name:  "Vase",
			Slug:  "vase",
			Image: "https://res.cloudinary.com/ichtrojan/image/upload/v1577751583/fragrance/perfume3.png",
			BottleSizes: []models.BottleSize{
				{
					Size:  "30 ml",
					Price: 4.0,
				},
				{
					Size:  "60 ml",
					Price: 8.0,
				},
				{
					Size:  "90 ml",
					Price: 12.5,
				},
			},
		},
		{
			Name:  "Mystic",
			Slug:  "mystic",
			Image: "https://res.cloudinary.com/ichtrojan/image/upload/v1577751580/fragrance/perfume2.png",
			BottleSizes: []models.BottleSize{
				{
					Size:  "30 ml",
					Price: 4.0,
				},
				{
					Size:  "60 ml",
					Price: 8.0,
				},
				{
					Size:  "90 ml",
					Price: 12.5,
				},
			},
		},
		{
			Name:  "Bold",
			Slug:  "bold",
			Image: "https://res.cloudinary.com/ichtrojan/image/upload/v1577751583/fragrance/perfume1.png",
			BottleSizes: []models.BottleSize{
				{
					Size:  "30 ml",
					Price: 4.0,
				},
				{
					Size:  "60 ml",
					Price: 8.0,
				},
				{
					Size:  "90 ml",
					Price: 12.5,
				},
			},
		},
	}

	for _, bottle := range bottles {
		db.Create(&bottle)
	}
}

func hashAndSalt(password []byte) string {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)

	if err != nil {
		log.Println(err)
	}

	return string(hash)
}
