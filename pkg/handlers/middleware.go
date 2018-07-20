package handlers

import (
	"github.com/gorilla/context"
	"github.com/roger-king/go-ecommerce/pkg/models"
	"github.com/roger-king/go-ecommerce/pkg/utilities"
	"github.com/sirupsen/logrus"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if len(token) > 0 {
			authedUser := models.Validate(models.JwtToken{Token: token})
			
			context.Set(r, 'user', authedUser)

			next.ServeHTTP(w, r)
		} else {
			utilities.RespondWithError(w, http.StatusUnauthorized, "not authorized")
		}
	})
}
