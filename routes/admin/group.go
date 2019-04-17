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
		r.Post("/{group_id}/users", controller.PostGroupUser)
		r.Ger("/{group_id}/users", controller.PostGroupUser) //TODO: create a controller to return all users from a group
		r.Delete("/{group_id}/users/{user_id}", controller.DeleteGroupUser)
		r.Post("/{group_id}/permissions", controller.PostGroupPermission)
		r.Get("/{group_id}/permissions", controller.PostGroupPermission) //TODO: create a controller to return all permissions from a group
		r.Delete("/{group_id}/permissions/{permission_id}", controller.DeleteGroupPermission)
	})

	return r
}
