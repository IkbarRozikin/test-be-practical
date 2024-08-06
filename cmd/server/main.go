package main

import (
	"log"
	"net/http"
	"test_be_practical/internal/routes"
	"test_be_practical/pkg/config"

	"github.com/gorilla/mux"
)

func main() {
	cfg := config.LoadConfig()
	router := mux.NewRouter()

	routes.RegisterRoutes(router)

	log.Println("Server running at port", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
		log.Fatal(err)
	}
}
