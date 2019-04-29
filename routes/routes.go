package routes

import (
	cryo "github.com/cryo-management/api/middlewares"
	"github.com/cryo-management/api/routes/admin"
	"github.com/cryo-management/api/routes/auth"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

// Setup configure the API endpoints
func Setup() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
		cryo.Authorization,
	)

	router.Route("/api/v1", func(r chi.Router) {
		r.Mount("/admin/configs", admin.ConfigRoutes())
		r.Mount("/admin/users", admin.UserRoutes())
		r.Mount("/admin/schemas", admin.SchemaRoutes())
		r.Mount("/admin/structures", admin.StructureRoutes())
		r.Mount("/admin/widgets", admin.WidgetRoutes())
		r.Mount("/admin/lookups", admin.LookupRoutes())
		r.Mount("/admin/groups", admin.GroupRoutes())
		r.Mount("/admin/currencies", admin.CurrencyRoutes())
		r.Mount("/auth", auth.Routes())
	})

	return router
}
