package admin

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"github.com/cryo-management/api/common"
	"github.com/cryo-management/api/models"

	"github.com/go-chi/render"
)

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

func AddUser(w http.ResponseWriter, r *http.Request) {
	GroupUser := new(models.GroupUser)
	body, err := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &GroupUser)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorParsingRequest, "AddUser unmarshal body", err.Error()))
		return
	}

	err = GroupUser.Create()
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorInsertingRecord, "AddUser creating", err.Error()))
		return
	}

	render.JSON(w, r, GroupUser)
}

func RemoveUser(w http.ResponseWriter, r *http.Request) {
	GroupUser := new(models.GroupUser)
	groupID := string(chi.URLParam(r, "group_id"))
	userID := string(chi.URLParam(r, "user_id"))

	GroupUser.GroupID = groupID
	GroupUser.UserID = userID

	err := GroupUser.Delete()
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorDeletingData, "RemoveUser delete group user", err.Error()))
		return
	}

	render.JSON(w, r, GroupUser)
}

func AddPermission(w http.ResponseWriter, r *http.Request) {
	GroupPermission := new(models.GroupPermission)
	body, err := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &GroupPermission)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorParsingRequest, "AddPermission unmarshal body", err.Error()))
		return
	}

	err = GroupPermission.Create()
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorInsertingRecord, "AddPermission creating", err.Error()))
		return
	}

	render.JSON(w, r, GroupPermission)
}

func RemovePermission(w http.ResponseWriter, r *http.Request) {
	GroupPermission := new(models.GroupPermission)
	groupID := string(chi.URLParam(r, "group_id"))
	structureID := string(chi.URLParam(r, "structure_id"))
	permissionType, err := strconv.Atoi(chi.URLParam(r, "type"))
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorDeletingData, "GroupPermission converting string to int", err.Error()))
		return
	}

	GroupPermission.GroupID = groupID
	GroupPermission.StructureID = structureID
	GroupPermission.Type = permissionType

	err = GroupPermission.Delete()
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, common.NewResponseError(common.ErrorDeletingData, "GroupPermission delete group user", err.Error()))
		return
	}

	render.JSON(w, r, GroupPermission)
}
