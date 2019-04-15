package routes

import (
	"github.com/cryo-management/api/routes/admin"
	myMiddlewares "github.com/cryo-management/api/routes/middlewares"
	"github.com/cryo-management/api/routes/structure"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

func Setup() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
		myMiddlewares.Session,
	)

	router.Route("/api/v1", func(r chi.Router) {
		r.Mount("/admin/users", admin.UserRoutes())
		r.Mount("/admin/schemas", admin.SchemaRoutes())
		r.Mount("/admin/schemas/{schema_id}/fields", admin.FieldRoutes())
		r.Mount("/admin/lookups", admin.LookupRoutes())
		r.Mount("/admin/lookups/{lookup_id}/lookups_options", admin.LookupOptionRoutes())
		r.Mount("/admin/groups", admin.GroupRoutes())
		r.Mount("/struct/schemas", structure.SchemaRoutes())
		//r.Mount("/api/data/{schema_id}", schema.Routes())
		//r.Mount("/api/auth", auth.Routes())
	})

	return router
}
