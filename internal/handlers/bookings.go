package handlers

import (
	"encoding/json"
	"net/http"
	"test_be_practical/internal/services"
)

func GetBookingsHandler(w http.ResponseWriter, r *http.Request) {
	bookings, err := services.GetBookings()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookings)
}
