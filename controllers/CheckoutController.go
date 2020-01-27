package controllers

import (
	"fmt"
	"net/http"
)

func ProcessCheckout(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, "thank you page")
}
