package controllers

import (
	"encoding/json"
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
	w.Header().Add("Content-Type", "application/json")

	creds := Credentials{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	db := database.Init()

	var admin models.Admin

	query := db.Where("email = ?", creds.Email).First(&admin)

	defer query.Close()

	if admin.Email == "" {
		_ = json.NewEncoder(w).Encode(Response{Message: "user not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(creds.Password)); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(Response{Message: "password not valid"})
		return
	}

	_ = json.NewEncoder(w).Encode(Response{Message: "login successful"})

	return
}

func Signin(w http.ResponseWriter, r *http.Request) {
	view = views.NewView("app", "signin")
	must(view.Render(w, nil))
}
