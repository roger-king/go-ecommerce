package handlers

import (
	"net/http"
	"github.com/roger-king/go-ecommerce/pkg/utilities"
	"github.com/roger-king/go-ecommerce/pkg/models"
	"encoding/json"
)

func AuthenticateController(w http.ResponseWriter, req *http.Request) {
	var authUser models.AuthUser

	decoder := json.NewDecoder(req.Body)

	if err := decoder.Decode(&authUser); err != nil {
		utilities.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer req.Body.Close()

	token := req.Header.Get("Authorization")

	if len(token) > 0 {
		authUser.Token = token
	}

	authedUser, err := models.Authenticate(authUser)

	if err != nil {
		utilities.RespondWithError(w, http.StatusBadRequest, "user already exists")
		return
	}

	utilities.RespondWithJSON(w, http.StatusCreated, authedUser)
}
