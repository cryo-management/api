package admin

import (
	"net/http"

	controller "github.com/cryo-management/api/controllers/admin"
	"github.com/go-chi/chi"
)

func LookupOptionRoutes() *chi.Mux {
	r := chi.NewRouter()

	// v1/api/admin/lookups-options
	r.Route("/", func(r chi.Router) {
		r.Post("/", controller.PostLookupOption)
		r.Get("/", controller.GetAllLookupOptions)
		r.Get("/{lookup_option_id}", controller.GetLookupOption)
		r.Patch("/{lookup_option_id}", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("update lookup_option"))
		})
		r.Delete("/{lookup_option_id}", controller.DeleteLookupOption)
	})

	return r
}
