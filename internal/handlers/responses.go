package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func RespondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")

	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("error marshaling JSON: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":"internal server error"}`))
		return
	}

	w.WriteHeader(status)
	w.Write(data)
}

func RespondWithError(w http.ResponseWriter, status int, msg string, err error) {
	if err != nil && status >= 500 {
		log.Printf("server error: %v", err)
	}

	RespondWithJSON(w, status, ErrorResponse{
		Error: msg,
	})
}
