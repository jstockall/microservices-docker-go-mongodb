package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mmorejon/cinema/common"
	"github.com/mmorejon/cinema/bookings/data"
	"github.com/mmorejon/cinema/bookings/models"
	"gopkg.in/mgo.v2"
)

// Handler for HTTP Get - "/health"
// Returns 200 if we can contact the DB
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	common.HealthCheck(w)
}

// Handler for HTTP Post - "/bookins"
// Create a new Booking document
func CreateBooking(w http.ResponseWriter, r *http.Request) {
	var dataResource BookingResource
	// Decode the incoming Booking json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid Booking data", 500)
		return
	}
	booking := &dataResource.Data
	// Create new context
	context := common.NewContext()
	defer context.Close()
	c := context.DbCollection("bookings")
	// Create Booking
	repo := &data.BookingRepository{c}
	repo.Create(booking)
	// Create response data
	j, err := json.Marshal(dataResource)
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func GetBookings(w http.ResponseWriter, r *http.Request) {
	// Create new context
	context := common.NewContext()
	defer context.Close()
	c := context.DbCollection("bookings")
	repo := &data.BookingRepository{c}

	query := r.URL.Query()
	var bookings []models.Booking
	if len(query["user"]) != 0 {
		bookings = repo.GetBy("userid", query["user"][0])
	} else if len(query["movie"]) != 0 {
		bookings = repo.GetBy("movieid", query["movie"][0])
	} else if len(query["showtime"]) != 0 {
		bookings = repo.GetBy("showtimeid", query["showtime"][0])
	} else {
		bookings = repo.GetAll()
	}

	// Create response data
	j, err := json.Marshal(BookingsResource{Data: bookings})
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// Handler for HTTP Get - "/bookings/{id}"
// Get movie by id
func GetBookingById(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	id := vars["id"]

	// create new context
	context := common.NewContext()
	defer context.Close()
	c := context.DbCollection("bookings")
	repo := &data.BookingRepository{c}

	// Get booking by id
	booking, err := repo.GetById(id)
	if err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		} else {
			common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
			return
		}
	}

	j, err := json.Marshal(BookingResource{Data: booking})
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// Handler for HTTP Delete - "/bookings/{id}"
// Delete a Booking document by id
func DeleteBooking(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	id := vars["id"]

	// Create new context
	context := common.NewContext()
	defer context.Close()
	c := context.DbCollection("bookings")

	// Remove showtime by id
	repo := &data.BookingRepository{c}
	err := repo.Delete(id)
	if err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		} else {
			common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
			return
		}
	}

	// Send response back
	w.WriteHeader(http.StatusNoContent)
}
