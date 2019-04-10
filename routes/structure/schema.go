package structure

import (
	controller "github.com/cryo-management/api/controllers/structure"
	"github.com/go-chi/chi"
)

//SchemaRoutes docs
func SchemaRoutes() *chi.Mux {
	r := chi.NewRouter()

	// v1/api/struct/schema/contract
	r.Route("/", func(r chi.Router) {
		r.Get("/{schema_code}", controller.GetSchema)
	})

	return r
}
