package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ichtrojan/fragrance/database"
	"github.com/ichtrojan/fragrance/models"

	"github.com/ichtrojan/fragrance/views"
)

var view *views.View

type CategoriesData struct {
	Categories []models.Category
}

type ScentsData struct {
	models.Scent
	Category string
}

type BottleData struct {
	models.Bottle
	Category string
	Scent    string
}

type BottleSize struct {
	models.BottleSize
}

type Fragrance struct {
	Image    string
	Category string
	Scent    string
	Bottle   string
	Sizes    []models.BottleSize
	Price    float64
}

func Home(w http.ResponseWriter, r *http.Request) {
	view = views.NewView("app", "home")
	must(view.Render(w, nil))
}

func Category(w http.ResponseWriter, r *http.Request) {
	view = views.NewView("app", "fragrance")

	db := database.Init()

	var categories []models.Category

	query := db.Find(&categories)

	defer query.Close()

	data := CategoriesData{
		Categories: categories,
	}

	must(view.Render(w, data))
}

func Scent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	category := vars["category"]

	view = views.NewView("app", "scent")

	db := database.Init()
	var scents []models.Scent

	var data []ScentsData

	query := db.Find(&scents)

	defer query.Close()

	for _, scent := range scents {
		data = append(data, ScentsData{
			Scent:    scent,
			Category: category,
		})
	}

	must(view.Render(w, data))
}

func Perfume(w http.ResponseWriter, r *http.Request) {
	view = views.NewView("app", "perfume")

	vars := mux.Vars(r)

	category := vars["category"]

	scent := vars["scent"]

	db := database.Init()

	var data []BottleData

	var bottles []models.Bottle

	query := db.Find(&bottles)

	defer query.Close()

	for _, bottle := range bottles {
		data = append(data, BottleData{
			Bottle:   bottle,
			Category: category,
			Scent:    scent,
		})
	}

	must(view.Render(w, data))
}

func Checkout(w http.ResponseWriter, r *http.Request) {
	view = views.NewView("app", "checkout")

	vars := mux.Vars(r)

	category := vars["category"]

	scent := vars["scent"]

	bottle := vars["bottle"]

	var fragrance models.Bottle

	var fragraceScent models.Scent

	// var bottleSize models.BottleSize

	db := database.Init()

	query := db.Where("slug = ?", bottle).Preload("BottleSizes").First(&fragrance)

	query = db.Where("slug = ?", scent).First(&fragraceScent)

	defer query.Close()

	data := Fragrance{
		Image:    fragrance.Image,
		Category: category,
		Scent:    scent,
		Bottle:   bottle,
		Sizes:    fragrance.BottleSizes,
		Price:    fragraceScent.Price,
	}
	// fmt.Printf("%+v\n", data)
	must(view.Render(w, data))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
