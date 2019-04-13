package admin

import (
	"net/http"

	controller "github.com/cryo-management/api/controllers/admin"
	"github.com/go-chi/chi"
)

//GroupRoutes docs
func GroupRoutes() *chi.Mux {
	r := chi.NewRouter()

	// v1/api/admin/group
	r.Route("/", func(r chi.Router) {
		r.Post("/", controller.PostGroup)
		r.Post("/{group_id}/users", controller.AddUser)
		r.Delete("/{group_id}/users/{user_id}", controller.RemoveUser)
		r.Get("/", controller.GetAllGroups)
		r.Get("/{group_id}", controller.GetGroup)
		r.Patch("/{group_id}", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("update group"))
		})
		r.Delete("/{group_id}", controller.DeleteGroup)
	})

	return r
}
