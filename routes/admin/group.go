package admin

import (
	"net/http"

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
		r.Patch("/{group_id}", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("update group"))
		})
		r.Delete("/{group_id}", controller.DeleteGroup)
		r.Post("/users", controller.AddUser)
		r.Delete("/{group_id}/users/{user_id}", controller.RemoveUser)
		r.Post("/permissions", controller.AddPermission)
		r.Delete("/{group_id}/permissions/{type}/structureid/{structure_id}", controller.RemovePermission)
	})

	return r
}
