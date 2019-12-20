package controllers

import (
	"html/template"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	view, err := template.ParseFiles("views/index.html")

	if err != nil {
		log.Fatal(nil)
	}

	err = view.Execute(w, nil)

	if err != nil {
		log.Fatal(nil)
	}
}
