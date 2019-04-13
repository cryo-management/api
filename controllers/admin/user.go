package admin

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/cryo-management/api/common"
	"github.com/cryo-management/api/models"

	"github.com/go-chi/render"
)

//PostUser docs
func PostUser(w http.ResponseWriter, r *http.Request) {
	user := new(models.User)
	body, err := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &user)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorParsingRequest, "PostUser unmarshal body", err.Error()))
		return
	}

	err = user.Create()
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorInsertingRecord, "PostUser creating", err.Error()))
		return
	}

	translation := new(models.Translation)
	err = translation.Create(user.ID, *user)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorInsertingRecord, "PostUser translation", err.Error()))
		return
	}

	render.JSON(w, r, user)
}

//GetUser docs
func GetUser(w http.ResponseWriter, r *http.Request) {
	user := new(models.User)
	id := string(chi.URLParam(r, "user_id"))

	err := user.Load(id)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorReturningData, "GetUser", err.Error()))
		return
	}

	render.JSON(w, r, user)
}

//GetAllUsers docs
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := new(models.Users)

	err := users.Load()
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorReturningData, "GetAllUsers load users", err.Error()))
		return
	}

	render.JSON(w, r, users)
}

//DeleteUser docs
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	user := new(models.User)
	id := string(chi.URLParam(r, "user_id"))

	err := user.Delete(id)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorDeletingData, "DeleteUser delete user", err.Error()))
		return
	}

	translation := new(models.Translation)

	err = translation.DeleteByStructureID(id)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorDeletingData, "DeleteUser delete translation", err.Error()))
		return
	}

	render.JSON(w, r, id)
}
