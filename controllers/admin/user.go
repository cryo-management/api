package admin

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/cryo-management/api/common"
	"github.com/cryo-management/api/models"
	services "github.com/cryo-management/api/services/admin"

	"github.com/go-chi/render"
)

func PostUser(w http.ResponseWriter, r *http.Request) {
	user := new(models.User)
	body, err := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &user)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorParsingRequest, "PostUser unmarshal body", err.Error()))
		return
	}

	userService := new(services.UserService)
	err = userService.Create(user)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorInsertingRecord, "PostUser creating", err.Error()))
		return
	}

	render.JSON(w, r, user)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	user := new(models.User)
	id := string(chi.URLParam(r, "user_id"))

	userService := new(services.UserService)
	err := userService.Load(user, id)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorReturningData, "GetUser load", err.Error()))
		return
	}

	render.JSON(w, r, user)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := new(models.Users)

	userService := new(services.UserService)
	err := userService.LoadAll(users)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorReturningData, "GetAllUsers load", err.Error()))
		return
	}

	render.JSON(w, r, users)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	user := new(models.User)
	id := string(chi.URLParam(r, "user_id"))

	userService := new(services.UserService)
	err := userService.Delete(user, id)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorDeletingData, "DeleteUser delete user", err.Error()))
		return
	}

	render.JSON(w, r, id)
}
