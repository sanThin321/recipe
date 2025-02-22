package httpresponse

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	// convert to json response
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
func HashPassword(password string) (string, error) {
	// Generate a bcrypt hash of the password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	// Convert the byte slice to a string and return it
	return string(hash), nil
}