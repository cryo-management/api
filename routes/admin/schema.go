package admin

import (
	controller "github.com/cryo-management/api/controllers/admin"
	"github.com/go-chi/chi"
)

func SchemaRoutes() *chi.Mux {
	r := chi.NewRouter()

	// v1/api/schema/admin/contract
	r.Route("/", func(r chi.Router) {
		r.Post("/", controller.PostSchema)
		// r.Get("/", controller.GetAllSchemas)
		// r.Get("/{schema_id}", controller.GetSchema)
		// r.Patch("/{schema_id}", func(w http.ResponseWriter, r *http.Request) {
		// 	w.Write([]byte("update schema"))
		// })
		// r.Delete("/{schema_id}", controller.DeleteSchema)
	})

	return r
}
