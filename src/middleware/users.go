package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/mbaezahuenupil/go-mongodb-test/src/models"
)

func UpdateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userUpdate := models.UserNew{}
		json.NewDecoder(r.Body).Decode(&userUpdate)
		validate := validator.New()
		err := validate.Struct(userUpdate)
		if err != nil {
			validationErrors := err.(validator.ValidationErrors)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			responseBody := map[string]string{"error": validationErrors.Error()}
			json.NewEncoder(w).Encode(responseBody)
			return
		}
		next.ServeHTTP(w, r)
	})
}
