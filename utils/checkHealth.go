package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status string `json:"status"`
}

func CheckHealth(w http.ResponseWriter, r *http.Request) {
	response := Response{Status: "Ok"}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
