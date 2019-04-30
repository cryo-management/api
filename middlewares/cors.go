package middlewares

import (
	"net/http"

	"github.com/go-chi/cors"
)

func Cors() *cors.Cors {
	cors := cors.New(cors.Options{
		AllowOriginFunc:  AllowOriginFunc,
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	return cors
}

func AllowOriginFunc(r *http.Request, origin string) bool {
	if origin == "http://localhost:8080" {
		return true
	}

	return false
}
