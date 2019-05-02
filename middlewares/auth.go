package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/cryo-management/api/models"
	"github.com/dgrijalva/jwt-go"
)

// Authorization validates the token and insert user data in the request
func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//  TODO: Lembrar de religar a verificação de autenticação
		if !strings.Contains(r.RequestURI, "/auth/login") && r.Method != "OPTIONS" {
			token, err := jwt.ParseWithClaims(r.Header.Get("Authorization"), &models.UserCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte("AllYourBase"), nil // TODO: Check the best place for this key, probably the config.toml
			})
			if token != nil && token.Valid {
				claims := token.Claims.(*models.UserCustomClaims)
				r.Header.Add("userID", claims.User.ID)
				r.Header.Add("Content-Language", claims.User.LanguageCode)
			} else {
				fmt.Println(err)
				http.Error(w, http.StatusText(401), http.StatusUnauthorized)
				return
			}
		}
		// r.Header.Add("userID", "307e481c-69c5-11e9-96a0-06ea2c43bb20")
		// r.Header.Add("Content-Language", "pt-br")

		next.ServeHTTP(w, r)
	})
}
