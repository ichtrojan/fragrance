package controllers

import (
	"github.com/ichtrojan/fragrance/database"
	"github.com/ichtrojan/fragrance/models"
	"net/http"

	"github.com/ichtrojan/fragrance/views"
)

var view *views.View

type CategoriesData struct {
	Categories []models.Category
}

func Home(w http.ResponseWriter, r *http.Request) {
	view = views.NewView("app", "home")
	must(view.Render(w, nil))
}

func Fragrance(w http.ResponseWriter, r *http.Request) {
	view = views.NewView("app", "fragrance")

	db := database.Init()

	var categories []models.Category

	db.Find(&categories)

	data := CategoriesData{
		Categories: categories,
	}

	must(view.Render(w, data))
}

func Scent(w http.ResponseWriter, r *http.Request) {
	view = views.NewView("app", "scent")
	must(view.Render(w, nil))
}

func Perfume(w http.ResponseWriter, r *http.Request) {
	view = views.NewView("app", "perfume")
	must(view.Render(w, nil))
}

func Checkout(w http.ResponseWriter, r *http.Request) {
	view = views.NewView("app", "checkout")
	must(view.Render(w, nil))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
