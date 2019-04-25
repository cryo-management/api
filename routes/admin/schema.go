package admin

import (
	controller "github.com/cryo-management/api/controllers/admin"
	"github.com/go-chi/chi"
)

// SchemaRoutes creates the api methods
func SchemaRoutes() *chi.Mux {
	r := chi.NewRouter()

	// v1/api/admin/schemas
	r.Route("/", func(r chi.Router) {
		r.Post("/", controller.PostSchema)
		r.Get("/", controller.GetAllSchemas)
		r.Get("/{schema_id}", controller.GetSchema)
		r.Patch("/{schema_id}", controller.UpdateSchema)
		r.Delete("/{schema_id}", controller.DeleteSchema)
		r.Post("/{schema_id}/fields", controller.PostField)
		r.Get("/{schema_id}/fields", controller.GetAllFields)
		r.Get("/{schema_id}/fields/{field_id}", controller.GetField)
		r.Patch("/{schema_id}/fields/{field_id}", controller.UpdateField)
		r.Delete("/{schema_id}/fields/{field_id}", controller.DeleteField)
		r.Post("/{schema_id}/pages", controller.PostPage)
		r.Get("/{schema_id}/pages", controller.GetAllPages)
		r.Get("/{schema_id}/pages/{page_id}", controller.GetPage)
		r.Patch("/{schema_id}/pages/{page_id}", controller.UpdatePage)
		r.Delete("/{schema_id}/pages/{page_id}", controller.DeletePage)
	})

	return r
}
