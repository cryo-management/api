package structure

import (
	"github.com/go-chi/chi"
)

func SchemaRoutes() *chi.Mux {
	r := chi.NewRouter()

	// v1/api/struct/schema/contract
	r.Route("/", func(r chi.Router) {
		// r.Get("/{schema_id}", controller.GetSchema)
	})

	return r
}
