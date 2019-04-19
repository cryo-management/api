package admin

import (
	controller "github.com/cryo-management/api/controllers/admin"
	"github.com/go-chi/chi"
)

// UserRoutes creates the api methods
func UserRoutes() *chi.Mux {
	r := chi.NewRouter()

	//  v1/api/admin/user
	r.Route("/", func(r chi.Router) {
		r.Post("/", controller.PostUser)
		r.Get("/", controller.GetAllUsers)
		r.Get("/{user_id}", controller.GetUser)
		r.Get("/{user_id}/groups", controller.GetAllGroupsByUser)
		r.Patch("/{user_id}", controller.UpdateUser)
		r.Delete("/{user_id}", controller.DeleteUser)
	})

	return r
}
