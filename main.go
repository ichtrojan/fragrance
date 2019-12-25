package main

import (
	"github.com/ichtrojan/fragrance/routes"
	"log"
	"net/http"
)

func main() {
	if err := http.ListenAndServe(":9000", routes.Init()); err != nil {
		log.Fatal(err)
	}
}
