package middlewares

import (
	"net/http"
)

//Session include informations into the request
func Session(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Header.Add("userID", "059fa339-025c-4104-ab55-c764d3028f63")
		r.Header.Add("firstName", "Bruno")
		r.Header.Add("lastName", "Piaui")
		r.Header.Add("email", "brunopiaui@gmail.com")
		r.Header.Add("languageCode", "pt-br")
		next.ServeHTTP(w, r)
	})
}
