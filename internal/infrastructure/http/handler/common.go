package handler

import (
	"encoding/json"
	"net/http"
)

type httpMethod string

const (
	GET    httpMethod = "GET"
	POST   httpMethod = "POST"
	DELETE httpMethod = "DELETE"
	PUT    httpMethod = "PUT"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func returnSuccess(w http.ResponseWriter, response any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func returnError(w http.ResponseWriter, message string, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	response := ErrorResponse{Message: message}
	json.NewEncoder(w).Encode(response)
}

func checkMethod(w http.ResponseWriter, r *http.Request, method httpMethod) bool {
	if r.Method != string(method) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return false
	}
	return true
}
