package controller

import (
	"net/http"

	"github.com/Phati/golang-boilerplate/middlewares"
	"github.com/gorilla/mux"
)

func InitRouter() (router *mux.Router) {
	router = mux.NewRouter()
	router.HandleFunc("/", middlewares.SetMiddlewareJSON(HomeController())).Methods(http.MethodGet)
	return
}
