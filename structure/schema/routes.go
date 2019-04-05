package schema

import (
	"github.com/go-chi/chi"
)

//Routes docs
func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", postSchema)
	router.Post("/", postSchema)
	router.Patch("/", postSchema)
	router.Delete("/", postSchema)
	return router
}
