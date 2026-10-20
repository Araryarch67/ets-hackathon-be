package lib

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}

func ReadJSON(r *http.Request, value any) error {
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(value); err != nil {
		return err
	}

	if err := decoder.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
		return errors.New("body hanya boleh berisi satu objek JSON")
	}

	return nil
}
