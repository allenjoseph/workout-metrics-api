package util

import (
	"encoding/json"
	"net/http"
)

// RespondError func
func RespondError(w http.ResponseWriter, code int, message string) {
	RespondJSON(w, code, message)
}

// RespondJSON func
func RespondJSON(w http.ResponseWriter, status int, payload interface{}) {
	res, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(res))
}
