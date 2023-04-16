package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wchopite/api-spark-v1/database"
	"github.com/wchopite/api-spark-v1/handlers"
	"github.com/wchopite/api-spark-v1/config"
)

func main() {
	config.Init()
	database.Init()
	c := config.GetConfig()
	router := mux.NewRouter()

	router.HandleFunc("/api/user", handlers.GetUser).Methods("GET")

	log.Fatal(http.ListenAndServe(":"+c.Port, router))
}
