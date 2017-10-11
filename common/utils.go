package common

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"fmt"
)

type (
	appError struct {
		Error      string `json:"error"`
		Message    string `json:"message"`
		HttpStatus int    `json:"status"`
	}
	errorResource struct {
		Data appError `json:"data"`
	}
	configuration struct {
		Server, MongoDBHost, DBUser, DBPwd, Database string
	}
)

// Handler for HTTP Get - "/health"
// Returns 200 if we can contact the DB
func HealthCheck(w http.ResponseWriter) {
	// Create new context
	context := NewContext()
	defer context.Close()
	err := context.Ping()

	var status []byte
	if err != nil {
		status = []byte(fmt.Sprintf(`{"status": "DOWN", "reason": "%s"}`, err))
	} else {
		status = []byte(`{"status": "UP"}`)
	}

	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(status)
}

func DisplayAppError(w http.ResponseWriter, handlerError error, message string, code int) {
	errObj := appError{
		Error:      handlerError.Error(),
		Message:    message,
		HttpStatus: code,
	}
	log.Printf("[AppError]: %s\n", handlerError)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if j, err := json.Marshal(errorResource{Data: errObj}); err == nil {
		w.Write(j)
	}
}


// AppConfig holds the configuration values from config.json file
var AppConfig configuration

// Initialize AppConfig
func initConfig() {
	file, err := os.Open("common/config.json")
	defer file.Close()
	if err != nil {
		log.Fatalf("[loadConfig]: %s\n", err)
	}
	decoder := json.NewDecoder(file)
	AppConfig = configuration{}
	err = decoder.Decode(&AppConfig)
	if err != nil {
		log.Fatalf("[logAppConfig]: %s\n", err)
	}
	dbHost := os.Getenv("DATABASE_HOST")
	if dbHost != "" {
		log.Printf("Setting DB host to " + dbHost)
		AppConfig.MongoDBHost = dbHost
	}
}
