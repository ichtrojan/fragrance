package routes

import (
	"github.com/gorilla/mux"
	"github.com/ichtrojan/fragrance/controllers"
	"html/template"
	"log"
	"net/http"
)

func Init() *mux.Router {
	route := mux.NewRouter()
	route.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./public"))))
	route.NotFoundHandler = http.HandlerFunc(notFound)
	route.HandleFunc("/", controllers.Home)

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
