package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mmorejon/cinema/common"
	"github.com/mmorejon/cinema/showtimes/data"
	"github.com/mmorejon/cinema/showtimes/models"
	"gopkg.in/mgo.v2"
)

// Handler for HTTP Get - "/health"
// Returns 200 if we can contact the DB
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	common.HealthCheck(w)
}

// Handler for HTTP Get - "/showtimes"
// Returns all Showtime documents
func GetShowTimes(w http.ResponseWriter, r *http.Request) {
	// Create new context
	context := common.NewContext()
	defer context.Close()
	c := context.DbCollection("showtimes")
	repo := &data.ShowTimeRepository{c}

	query := r.URL.Query()
	movie := query["movie"]
	var showtimes []models.ShowTime
	if len(movie) == 0 {
		// Get all showtimes form repository
		showtimes = repo.GetAll()
	} else {
		// Filter by movie
		showtimes = repo.GetByMovie(movie[0])
	}

	j, err := json.Marshal(ShowTimesResource{Data: showtimes})
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// Handler for HTTP Post - "/showtimes"
// Create a new Showtime document
func CreateShowTime(w http.ResponseWriter, r *http.Request) {
	var dataResource ShowTimeResource
	// Decode the incoming ShowTime json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid ShowTime data", 500)
		return
	}
	showtime := &dataResource.Data
	// Create new context
	context := common.NewContext()
	defer context.Close()
	c := context.DbCollection("showtimes")
	// Create ShowTime
	repo := &data.ShowTimeRepository{c}
	repo.Create(showtime)
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

func GetShowTimeById(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	id := vars["id"]

	// Create new context
	context := common.NewContext()
	defer context.Close()
	c := context.DbCollection("showtimes")
	repo := &data.ShowTimeRepository{c}

	// Get showtime by id
	showtime, err := repo.GetById(id)
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	// Create data for the response
	j, err := json.Marshal(ShowTimeResource{Data: showtime})
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// Handler for HTTP Delete - "/showtimes/{id}"
// Delete a Showtime document by id
func DeleteShowTime(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	id := vars["id"]

	// Create new context
	context := common.NewContext()
	defer context.Close()
	c := context.DbCollection("showtimes")

	// Remove showtime by id
	repo := &data.ShowTimeRepository{c}
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

// Handler for HTTP Put - "/showtimes/{id}"
// Update a Showtime document
func UpdateShowTime(w http.ResponseWriter, r *http.Request) {
	var dataResource ShowTimeResource
	// Decode the incoming ShowTime json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid ShowTime data", 500)
		return
	}

	showtime := &dataResource.Data
	// Create new context
	context := common.NewContext()
	defer context.Close()
	c := context.DbCollection("showtimes")
	// Create ShowTime
	repo := &data.ShowTimeRepository{c}
	repo.Update(showtime)
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
