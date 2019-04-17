package admin

import (
	controller "github.com/cryo-management/api/controllers/admin"
	"github.com/go-chi/chi"
)

func LookupRoutes() *chi.Mux {
	r := chi.NewRouter()

	// v1/api/admin/lookups
	r.Route("/", func(r chi.Router) {
		r.Post("/", controller.PostLookup)
		r.Get("/", controller.GetAllLookups)
		r.Get("/{lookup_id}", controller.GetLookup)
		r.Patch("/{lookup_id}", controller.UpdateLookup)
		r.Delete("/{lookup_id}", controller.DeleteLookup)
	})

	return r
}
