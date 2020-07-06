package util

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, httStatus int, payload interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httStatus)
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	if _, err := w.Write(b); err != nil {
		return err
	}
	return nil
}
