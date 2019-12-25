package controllers

import (
	"github.com/ichtrojan/fragrance/views"
	"net/http"
)

var view *views.View

func Home(w http.ResponseWriter, r *http.Request) {
	view = views.NewView("app", "home")
	must(view.Render(w, nil))
}

func Fragrance(w http.ResponseWriter, r *http.Request) {
	view = views.NewView("app", "fragrance")
	must(view.Render(w, nil))
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
