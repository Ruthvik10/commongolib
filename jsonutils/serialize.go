package jsonutils

import (
	"encoding/json"
	"net/http"
)

func ToJSON(w http.ResponseWriter, status int, data any, headers http.Header) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}
	for k, v := range headers {
		w.Header()[k] = v
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		return err
	}
	return nil
}
