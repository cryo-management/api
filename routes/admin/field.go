package admin

import (
	controller "github.com/cryo-management/api/controllers/admin"
	"github.com/go-chi/chi"
)

func FieldRoutes() *chi.Mux {
	r := chi.NewRouter()

	// v1/api/admin/contract/field
	r.Route("/", func(r chi.Router) {
		r.Post("/", controller.PostField)
		r.Get("/", controller.GetAllFields)
		r.Get("/{field_id}", controller.GetField)
		r.Patch("/{field_id}", controller.UpdateField)
		r.Delete("/{field_id}", controller.DeleteField)
	})

	return r
}
