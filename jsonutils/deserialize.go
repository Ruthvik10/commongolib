package jsonutils

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func ReadJSON(w http.ResponseWriter, r *http.Request, dest any) error {
	const maxBytes = 1_048_576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(dest)
	if err != nil {
		return err
	}
	err = dec.Decode(&struct{}{})
	if !errors.Is(err, io.EOF) {
		return errors.New("body must contain a single JSON value")
	}
	return nil
}
