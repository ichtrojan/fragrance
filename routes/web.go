package routes

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ichtrojan/fragrance/controllers"
)

func Init() *mux.Router {
	route := mux.NewRouter()
	route.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./public"))))
	route.NotFoundHandler = http.HandlerFunc(notFound)

	route.HandleFunc("/", controllers.Home)
	route.HandleFunc("/categories", controllers.Category)
	route.HandleFunc("/{category}", controllers.Scent)
	route.HandleFunc("/{category}/{scent}", controllers.Perfume)
	route.HandleFunc("/{category}/{scent}/{bottle}/checkout", controllers.Checkout)
	route.HandleFunc("/login", controllers.LoginUser).Methods("POST")

	return route
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)

	view, err := template.ParseFiles("views/errors/404.html")

	if err != nil {
		log.Fatal(err)
	}

	_ = view.Execute(w, nil)
}
