package schema

import (
	"github.com/go-chi/chi"
)

//Routes docs
func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Post("/", postSchema)
	router.Get("/", getAllSchemas)
	return router
}
