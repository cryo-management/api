package routes

import (
	cryoMiddleware "github.com/cryo-management/api/middlewares"
	"github.com/cryo-management/api/routes/admin"
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
		cryoMiddleware.Session,
	)

	router.Route("/api/v1", func(r chi.Router) {
		r.Mount("/admin/users", admin.UserRoutes())
		r.Mount("/admin/schemas", admin.SchemaRoutes())
		r.Mount("/admin/lookups", admin.LookupRoutes())
		r.Mount("/admin/groups", admin.GroupRoutes())
	})

	return router
}
