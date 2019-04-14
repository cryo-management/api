package middlewares

import (
	"net/http"

	"github.com/cryo-management/api/common"
)

func Session(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		common.Session.User.ID = "059fa339-025c-4104-ab55-c764d3028f63"
		common.Session.User.FirstName = "Bruno"
		common.Session.User.LastName = "Piaui"
		common.Session.User.Email = "brunopiaui@gmail.com"
		common.Session.User.Language = "pt-br"
		next.ServeHTTP(w, r)
	})
}
