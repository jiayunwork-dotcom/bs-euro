package api

import (
	"encoding/json"
	"net/http"

	"bs-euro/internal/model"
)

// ErrorBody is the stable error shape.
type ErrorBody struct {
	Error ErrorDetail `json:"error"`
}

// ErrorDetail carries the machine code and message.
type ErrorDetail struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func writeError(w http.ResponseWriter, err error) {
	body := ErrorBody{
		Error: ErrorDetail{
			Code:    string(model.Code(err)),
			Message: err.Error(),
		},
	}
	writeJSON(w, http.StatusBadRequest, body)
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func methodNotAllowed(w http.ResponseWriter) {
	writeError(w, model.NewError(model.CodeUnknownCommand, "method not allowed"))
}

func notFound(w http.ResponseWriter) {
	writeError(w, model.NewError(model.CodeUnknownCommand, "route not found"))
}
