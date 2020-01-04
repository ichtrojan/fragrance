package controllers

import (
	"github.com/gorilla/securecookie"
	"golang.org/x/crypto/bcrypt"
	"net/http"

	"github.com/ichtrojan/fragrance/database"
	"github.com/ichtrojan/fragrance/models"
	"github.com/ichtrojan/fragrance/views"
)

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

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

	setSession(w, admin.ID)

	http.Redirect(w, r, "dashboard", 302)
}

func Signin(w http.ResponseWriter, r *http.Request) {
	view = views.NewView("app", "signin")
	must(view.Render(w, nil))
}

func setSession(w http.ResponseWriter, userId uint) {
	value := map[string]uint{
		"id": userId,
	}

	if encoded, err := cookieHandler.Encode("gofrag", value); err == nil {
		cookie := &http.Cookie{
			Name:   "gofrag",
			Value:  encoded,
			Path:   "/",
			MaxAge: 3600,
		}

		http.SetCookie(w, cookie)
	}
}

func GetSession(w http.ResponseWriter, r *http.Request) (id uint) {
	if cookie, err := r.Cookie("gofrag"); err == nil {
		cookieValue := make(map[string]uint)

		if err = cookieHandler.Decode("gofrag", cookie.Value, &cookieValue); err == nil {
			id = cookieValue["id"]
			return id
		}
	}

	http.Redirect(w, r, "login", 302)
	return
}

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:   "gofrag",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}

	http.SetCookie(w, cookie)

	http.Redirect(w, r, "login", 302)
}
