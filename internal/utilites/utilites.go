package utilites

import (
	"encoding/json"
	"net/http"
	"unicode"
)

func DecodeJsonBody(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&dst)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return err
	}
	return nil
}

func ResponseWithJson(w http.ResponseWriter, r *http.Request, code int, payload interface{}) error {
	data, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return err
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(data)
	return nil
}

func ResponseWithError(w http.ResponseWriter, r *http.Request, code int, message string) {
	type Error struct {
		Error string `json:"error"`
	}
	errResponse := Error{Error: message}
	ResponseWithJson(w, r, code, errResponse)

}

func IsValidUsername(text string) bool {
	if len(text) < 3 {
		return false
	}

	for _, char := range text {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
			return false
		}
	}
	return true

}
