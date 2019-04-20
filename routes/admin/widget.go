package admin

import (
	controller "github.com/cryo-management/api/controllers/admin"
	"github.com/go-chi/chi"
)

// WidgetRoutes creates the api methods
func WidgetRoutes() *chi.Mux {
	r := chi.NewRouter()

	// v1/api/admin/widgets
	r.Route("/", func(r chi.Router) {
		r.Post("/", controller.PostWidget)
		r.Get("/", controller.GetAllWidgets)
		r.Get("/{schema_id}", controller.GetWidget)
		r.Patch("/{schema_id}", controller.UpdateWidget)
		r.Delete("/{schema_id}", controller.DeleteWidget)
	})

	return r
}
