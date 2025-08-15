package http

import (
	"encoding/json"
	"net/http"
)

type apiError struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
}

func JSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func Error(w http.ResponseWriter, status int, err error, msg string) {
	JSON(w, status, apiError{
		Error:   http.StatusText(status),
		Message: msgOrErr(msg, err),
	})
}

func msgOrErr(msg string, err error) string {
	if msg != "" {
		return msg
	}
	if err != nil {
		return err.Error()
	}
	return ""
}
