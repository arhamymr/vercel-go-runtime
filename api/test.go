package handler

import (
	"encoding/json"
	"net/http"

	"rsc.io/quote"
)
 
type Response struct {
	Message string `json:"message"`
	Quote   string `json:"quote"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Message: "Hello from Go!",
		Quote:   quote.Go(),
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}