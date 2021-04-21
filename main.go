package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Phati/golang-boilerplate/config"
	"github.com/Phati/golang-boilerplate/controller"
	"github.com/Phati/golang-boilerplate/db"
)

func main() {
	config.InitConfigs()
	db.InitDB()
	fmt.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":"+config.AppPort(), controller.InitRouter()))
}
