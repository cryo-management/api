package admin

import (
	"net/http"

	controller "github.com/cryo-management/api/controllers/admin"
	"github.com/go-chi/chi"
)

//SchemaRoutes docs
func SchemaRoutes() *chi.Mux {
	r := chi.NewRouter()

	// v1/api/schema/admin/contract
	r.Route("/", func(r chi.Router) {
		r.Post("/", controller.PostSchema)
		r.Get("/", controller.GetAllSchemas)
		r.Get("/{schema_code}", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("get schema instance"))
		})
		r.Patch("/{schema_code}", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("update schema"))
		})
		r.Delete("/{schema_code}", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("delete schema"))
		})
	})

	return r
}
