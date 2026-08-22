package util

import (
	"encoding/json"
	"net/http"
	"sync"
)

// M is a general-purpose map
// mainly to hold response payload.
type M map[string]any

// successPayload is the base structure of a success response.
type successPayload struct {
	StatusText string `json:"status_text"`
	StatusCode int    `json:"status_code"`
	Data       M      `json:"data"`
}

type successPayloadV2 struct {
	StatusText string `json:"status_text"`
	StatusCode int    `json:"status_code"`
	Data       any    `json:"data"`
}

// errorPayload is the base structure of an error response.
type errorPayload struct {
	StatusText string `json:"status_text"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

// Reusable success payload to reduce memory allocation.
var successPayloadPool = sync.Pool{
	New: func() any {
		return new(successPayloadV2)
	},
}

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

func SendJSON(w http.ResponseWriter, statusCode int, data any) {
	payload := successPayloadPool.Get().(*successPayloadV2)
	defer successPayloadPool.Put(payload)

	payload.StatusCode = statusCode
	payload.StatusText = http.StatusText(statusCode)
	payload.Data = data

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(payload)
}

// JSONError sets the `Content-Type` header to json.
// It does not carry payload. Instead, it sends error string as message.
func JSONError(w http.ResponseWriter, statusCode int, err error) {
	bytesJSON, _ := json.Marshal(errorPayload{
		StatusText: http.StatusText(statusCode),
		StatusCode: statusCode,
		Message:    err.Error(),
	})
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(bytesJSON)
}
