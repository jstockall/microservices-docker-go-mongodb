package controllers

import (
	"encoding/json"
	"net/http"
	"fmt"
	"log"
	"os/exec"

	"github.com/mmorejon/cinema/backup/common"
)

// Handler for HTTP Get - "/"
func Health(w http.ResponseWriter, r *http.Request) {
	// Create response data
	j, err := json.Marshal("Post to /backup or /restore")
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}

	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// Handler for HTTP Post - "/backup"
func Backup(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Begin Backup\n")
	out, err := exec.Command("sh", "-c", "./backup.sh").CombinedOutput()
	if err != nil {
		log.Print(err)
		common.DisplayAppError(w, err, "Unable to backup", 500)
		return
	}

	fmt.Printf("The Output is %s\n", string(out))

	// Create response data
	j, err := json.Marshal(string(out))
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}

	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// Handler for HTTP Post - "/restore"
func Restore(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Begin Restore\n")
	out, err := exec.Command("sh", "-c", "./restore.sh").CombinedOutput()
	if err != nil {
		log.Print(err)
		common.DisplayAppError(w, err, "Unable to restore", 500)
	}
	fmt.Printf("The Output is:\n%s", string(out))

	// Create response data
	j, err := json.Marshal(string(out))
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}

	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
