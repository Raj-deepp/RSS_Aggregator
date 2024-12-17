package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Add("Content-Type", "application/json")
	data, err := json.Marshal(payload)

	if err != nil {
		log.Printf("Failed to Marshal JSON response: %v", err)
		errorResponse := map[string]string{"error": "Internal Server Error"}
		errorData, _ := json.Marshal(errorResponse)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write(errorData)
		return
	}

	w.WriteHeader(code)
	w.Write(data)
}
