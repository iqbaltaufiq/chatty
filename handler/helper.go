package handler

import (
	"encoding/json"
	"net/http"
)

// JSONSuccess sets the `Content-Type` header to json before writing `data` into ResponseWriter.
func JSONSuccess(w http.ResponseWriter, statusCode int, data M) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	payload := successPayload{
		StatusText: http.StatusText(statusCode),
		StatusCode: statusCode,
		Data:       data,
	}
	json.NewEncoder(w).Encode(&payload)
}

// JSONError sets the `Content-Type` header to json.
// It does not carry payload. Instead, it sends error string as message.
func JSONError(w http.ResponseWriter, statusCode int, err error) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	payload := errorPayload{
		StatusText: http.StatusText(statusCode),
		StatusCode: statusCode,
		Message:    err.Error(),
	}
	json.NewEncoder(w).Encode(&payload)
}
