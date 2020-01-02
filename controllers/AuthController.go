package controllers

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/ichtrojan/fragrance/database"
	"github.com/ichtrojan/fragrance/models"
	"github.com/ichtrojan/fragrance/views"
)

type Response struct {
	Message string
}

type Credentials struct {
	Email    string
	Password string
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	creds := Credentials{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	db := database.Init()

	var admin models.Admin

	query := db.Where("email = ?", creds.Email).First(&admin)

	defer query.Close()

	if admin.Email == "" {
		data := Response{Message: "user not found"}
		view = views.NewView("app", "signin")
		must(view.Render(w, data))
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(creds.Password)); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		data := Response{Message: "password not valid"}
		view = views.NewView("app", "signin")
		must(view.Render(w, data))
		return
	}

	data := Response{Message: "login successful"}
	view = views.NewView("app", "signin")
	must(view.Render(w, data))
	return
}

func Signin(w http.ResponseWriter, r *http.Request) {
	view = views.NewView("app", "signin")
	must(view.Render(w, nil))
}
