package utils

import (
	"encoding/json"
	"net/http"
)



func ResponseHttpSuccess(w http.ResponseWriter, statusCode int, data interface{}) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": statusCode,
		"message": "Success",
		"data": data,
	})
}

func ResponseHttpError(w http.ResponseWriter, statusCode int, message string) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": statusCode,
		"message": message,
		"data": nil,
	})
}