package handlers

import (
	"github.com/roger-king/go-ecommerce/pkg/models"
	"github.com/roger-king/go-ecommerce/pkg/utilities"
	"github.com/sirupsen/logrus"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if len(token) > 0 {
			isAuthed := models.Validate(models.JwtToken{Token: token})

			logrus.Infoln("PLease work", isAuthed)

			next.ServeHTTP(w, r)
		} else {
			utilities.RespondWithError(w, http.StatusUnauthorized, "not authorized")
		}
	})
}
