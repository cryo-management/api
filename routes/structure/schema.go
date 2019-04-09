package structure

import (
	controller "github.com/cryo-management/api/controllers/structure"
	"github.com/go-chi/chi"
)

//SchemaRoutes docs
func SchemaRoutes() *chi.Mux {
	r := chi.NewRouter()

	// v1/api/schema/admin/contract
	r.Route("/", func(r chi.Router) {
		r.Get("/", controller.GetSchema)
	})

	return r
}
