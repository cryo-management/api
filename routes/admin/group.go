package admin

import (
	controller "github.com/cryo-management/api/controllers/admin"
	"github.com/go-chi/chi"
)

func GroupRoutes() *chi.Mux {
	r := chi.NewRouter()

	// v1/api/admin/group
	r.Route("/", func(r chi.Router) {
		r.Post("/", controller.PostGroup)
		r.Get("/", controller.GetAllGroups)
		r.Get("/{group_id}", controller.GetGroup)
		r.Patch("/{group_id}", controller.UpdateGroup)
		r.Delete("/{group_id}", controller.DeleteGroup)
		r.Post("/users", controller.PostGroupUser)
		r.Delete("/{group_id}/users/{user_id}", controller.DeleteGroupUser)
		r.Post("/permissions", controller.PostGroupPermission)
		r.Delete("/{group_id}/permissions/{type}/structure_id/{structure_id}", controller.DeleteGroupPermission) //TODO alterar p id da permissão
	})

	return r
}
