package routers

import (
	"github.com/gorilla/mux"
	"github.com/mmorejon/cinema/showtimes/controllers"
)

func SetShowTimeRouters(router *mux.Router) *mux.Router {
	router.HandleFunc("/health", controllers.HealthCheck).Methods("GET")
	router.HandleFunc("/showtimes", controllers.GetShowTimes).Methods("GET")
	router.HandleFunc("/showtimes", controllers.CreateShowTime).Methods("POST")
	router.HandleFunc("/showtimes/{id}", controllers.GetShowTimeById).Methods("GET")
	router.HandleFunc("/showtimes/{id}", controllers.DeleteShowTime).Methods("DELETE")
	router.HandleFunc("/showtimes/{id}", controllers.UpdateShowTime).Methods("PUT")
	return router
}
