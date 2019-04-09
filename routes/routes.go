package routes

import (
	"github.com/cryo-management/api/routes/admin"
	"github.com/cryo-management/api/routes/structure"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

//Setup docs
func Setup() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)

	router.Route("/api/v1", func(r chi.Router) {
		r.Mount("/admin/schema", admin.SchemaRoutes())
		r.Mount("/struct/schema", structure.SchemaRoutes())
		//r.Mount("/api/data/{schema_code}", schema.Routes())
		//r.Mount("/api/auth", auth.Routes())
	})

	return router
}
