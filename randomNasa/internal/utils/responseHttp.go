package utils

import (
	"encoding/json"
	"net/http"
)

func ResponsseSuccess(w http.ResponseWriter, statusRes int, data interface{})  {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": statusRes,
		"data": data,
	})
}

func ResponseError(w http.ResponseWriter, statusRes int, message string)  {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": statusRes,
		"message": message,
		"data": nil,
	})
}