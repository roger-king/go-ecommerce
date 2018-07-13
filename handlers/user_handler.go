package handlers

import (
	"net/http"
	"github.com/roger-king/go-ecommerce/models"
	"encoding/json"
	"github.com/roger-king/go-ecommerce/utilities"
)

func CreateUserController(w http.ResponseWriter, req *http.Request) {
	var u models.User

	decoder := json.NewDecoder(req.Body)

	if err := decoder.Decode(&u); err != nil {
		utilities.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer req.Body.Close()

	models.CreateUser(u)

	utilities.RespondWithJSON(w, http.StatusCreated, u)
}
