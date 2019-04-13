package admin

import (
	"net/http"

	controller "github.com/cryo-management/api/controllers/admin"
	"github.com/go-chi/chi"
)

//UserRoutes docs
func UserRoutes() *chi.Mux {
	r := chi.NewRouter()

	// v1/api/admin/user
	r.Route("/", func(r chi.Router) {
		r.Post("/", controller.PostUser)
		r.Get("/", controller.GetAllUsers)
		r.Get("/{user_id}", controller.GetUser)
		r.Patch("/{user_id}", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("update user"))
		})
		r.Delete("/{user_id}", controller.DeleteUser)
	})

	return r
}
