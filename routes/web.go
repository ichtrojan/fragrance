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
	route.HandleFunc("/fragrance", controllers.Fragrance)
	route.HandleFunc("/scent/{scent}", controllers.Scent)
	route.HandleFunc("/perfume/bottle", controllers.Perfume)
	route.HandleFunc("/checkout", controllers.Checkout)

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
