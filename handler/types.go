package handler

// M is a general-purpose map
// mainly to hold response payload.
type M map[string]any

// successPayload is the base structure of a success response.
type successPayload struct {
	StatusText string `json:"status_text"`
	StatusCode int    `json:"status_code"`
	Data       M      `json:"data"`
}

// errorPayload is the base structure of an error response.
type errorPayload struct {
	StatusText string `json:"status_text"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}
