package handlers

import (
	"net/http"
	"github.com/roger-king/go-ecommerce/pkg/models"
	"encoding/json"
	"github.com/roger-king/go-ecommerce/pkg/utilities"
)

func CreateUserController(w http.ResponseWriter, req *http.Request) {
	var u models.User

	decoder := json.NewDecoder(req.Body)

	if err := decoder.Decode(&u); err != nil {
		utilities.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer req.Body.Close()

	createdUser, createErr := models.CreateUser(u)

	if createErr != nil {
		utilities.RespondWithError(w, http.StatusBadRequest, "user already exists")
		return
	}

	utilities.RespondWithJSON(w, http.StatusCreated, createdUser)
}


func FindUserByEmailController(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()

	email, queryErr := query["email"]

	if !queryErr {
		utilities.RespondWithError(w, http.StatusBadRequest, "invalid request")
		return
	}

	user, err := models.FindUserByEmail(email[0])

	if err != nil {
		utilities.RespondWithError(w, http.StatusBadRequest, "user does not exist")
		return
	}

	utilities.RespondWithJSON(w, http.StatusFound, user)
}
