package handlers

import (
	"net/http"
	"github.com/roger-king/go-ecommerce/pkg/utilities"
	"github.com/roger-king/go-ecommerce/pkg/models"
	"encoding/json"
)

func AuthenticateController(w http.ResponseWriter, req *http.Request) {
	var user models.User

	decoder := json.NewDecoder(req.Body)

	if err := decoder.Decode(&user); err != nil {
		utilities.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer req.Body.Close()

	token, err := models.Authenticate(user)

	if err != nil {
		utilities.RespondWithError(w, http.StatusBadRequest, "Invalid email or password")
		return
	}

	utilities.RespondWithJSON(w, http.StatusCreated, token)
}
