package routers

import (
	"github.com/gorilla/mux"
	"github.com/mmorejon/cinema/users/controllers"
)

func SetUsersRouters(router *mux.Router) *mux.Router {
	router.HandleFunc("/health", controllers.HealthCheck).Methods("GET")
	router.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")
	router.HandleFunc("/users/{id}", controllers.GetUserById).Methods("GET")
	return router
}
