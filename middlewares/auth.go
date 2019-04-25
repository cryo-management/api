package middlewares

import (
	"net/http"
)

// Authorization validates the token and insert user data in the request
func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//  TODO: Lembrar de religar a verificação de autenticação
		//  if !strings.Contains(r.RequestURI, "/auth/login") {
		//  	token, err := jwt.ParseWithClaims(r.Header.Get("Authorization"), &models.UserCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		//  		return []byte("AllYourBase"), nil // TODO: Check the best place for this key, probably the config.toml
		//  	})
		//  	if token != nil && token.Valid {
		//  		claims := token.Claims.(*models.UserCustomClaims)
		//  		r.Header.Add("userID", claims.User.ID)
		//  		r.Header.Add("languageCode", claims.User.LanguageCode)
		//  	} else {
		//  		fmt.Println(err)
		//  		http.Error(w, http.StatusText(401), http.StatusUnauthorized)
		//  		return
		//  	}
		//  }
		r.Header.Add("userID", "57a97aaf-16da-44ef-a8be-b1caf52becd6")
		r.Header.Add("languageCode", "pt-br")

		next.ServeHTTP(w, r)
	})
}
