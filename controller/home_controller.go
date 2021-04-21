package controller

import (
	"net/http"

	"github.com/Phati/golang-boilerplate/responses"
)

func HomeController() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		responses.JSON(w, http.StatusAccepted, responses.MessageResponse{Message: "Hello from Homecontroller"})
	}
}
