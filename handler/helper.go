package handler

import (
	"encoding/json"
	"net/http"
)

// jsonSuccess sets the `Content-Type` header to json before writing `data` into ResponseWriter.
func jsonSuccess(w http.ResponseWriter, statusCode int, data M) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	payload := successPayload{
		StatusText: http.StatusText(statusCode),
		StatusCode: statusCode,
		Data:       data,
	}
	json.NewEncoder(w).Encode(&payload)
}

// jsonError sets the `Content-Type` header to json.
// It does not carry payload. Instead, it sends error string as message.
func jsonError(w http.ResponseWriter, statusCode int, err error) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	payload := errorPayload{
		StatusText: http.StatusText(statusCode),
		StatusCode: statusCode,
		Message:    err.Error(),
	}
	json.NewEncoder(w).Encode(&payload)
}
