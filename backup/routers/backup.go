package routers

import (
	"github.com/gorilla/mux"
	"github.com/mmorejon/cinema/backup/controllers"
)

func SetBackupRouters(router *mux.Router) *mux.Router {
	router.HandleFunc("/", controllers.Health).Methods("GET")
	router.HandleFunc("/backup", controllers.Backup).Methods("POST")
	router.HandleFunc("/restore", controllers.Restore).Methods("POST")
	return router
}
