package routes

import (
	"test_be_practical/internal/handlers"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/bookings", handlers.GetBookingsHandler).Methods("GET")
}
