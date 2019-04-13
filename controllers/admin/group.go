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

// PostGroup docs
func PostGroup(w http.ResponseWriter, r *http.Request) {
	group := new(models.Group)
	body, err := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &group)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorParsingRequest, "PostGroup unmarshal body", err.Error()))
		return
	}

	err = group.Create()
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorInsertingRecord, "PostGroup creating", err.Error()))
		return
	}

	translation := new(models.Translation)
	err = translation.Create(group.ID, *group)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorInsertingRecord, "PostGroup translation", err.Error()))
		return
	}

	render.JSON(w, r, group)
}

// AddUser docs
func AddUser(w http.ResponseWriter, r *http.Request) {
	groupsUsers := new(models.GroupsUsers)
	groupID := string(chi.URLParam(r, "group_id"))
	body, err := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &groupsUsers)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorParsingRequest, "AddUser unmarshal body", err.Error()))
		return
	}

	groupsUsers.GroupID = groupID

	err = groupsUsers.AddUser()
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorInsertingRecord, "AddUser creating", err.Error()))
		return
	}

	render.JSON(w, r, groupsUsers)
}

// RemoveUser docs
func RemoveUser(w http.ResponseWriter, r *http.Request) {
	groupsUsers := new(models.GroupsUsers)
	groupID := string(chi.URLParam(r, "group_id"))
	userID := string(chi.URLParam(r, "user_id"))

	groupsUsers.GroupID = groupID
	groupsUsers.UserID = userID

	err := groupsUsers.RemoveUser()
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorDeletingData, "RemoveUser delete group user", err.Error()))
		return
	}

	render.JSON(w, r, groupsUsers)
}

//GetGroup docs
func GetGroup(w http.ResponseWriter, r *http.Request) {
	group := new(models.Group)
	id := string(chi.URLParam(r, "group_id"))

	err := group.Load(id)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorReturningData, "GetGroup", err.Error()))
		return
	}

	render.JSON(w, r, group)
}

//GetAllGroups docs
func GetAllGroups(w http.ResponseWriter, r *http.Request) {
	groups := new(models.Groups)

	err := groups.Load()
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorReturningData, "GetAllGroups load groups", err.Error()))
		return
	}

	render.JSON(w, r, groups)
}

//DeleteGroup docs
func DeleteGroup(w http.ResponseWriter, r *http.Request) {
	group := new(models.Group)
	id := string(chi.URLParam(r, "group_id"))

	err := group.Delete(id)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorDeletingData, "DeleteGroup delete group", err.Error()))
		return
	}

	translation := new(models.Translation)

	err = translation.DeleteByStructureID(id)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorDeletingData, "DeleteGroup delete translation", err.Error()))
		return
	}

	render.JSON(w, r, id)
}
