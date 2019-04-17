package admin

import (
	"net/http"

	"github.com/cryo-management/api/services"

	"github.com/go-chi/render"
)

func PostGroup(w http.ResponseWriter, r *http.Request) {
	response := services.CreateGroup(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

func GetAllGroups(w http.ResponseWriter, r *http.Request) {
	response := services.LoadAllGroups(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

func GetGroup(w http.ResponseWriter, r *http.Request) {
	response := services.LoadGroup(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

func UpdateGroup(w http.ResponseWriter, r *http.Request) {
	response := services.UpdateGroup(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

func DeleteGroup(w http.ResponseWriter, r *http.Request) {
	response := services.DeleteGroup(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

func PostGroupUser(w http.ResponseWriter, r *http.Request) {
	response := services.CreateGroupUser(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

func DeleteGroupUser(w http.ResponseWriter, r *http.Request) {
	response := services.DeleteGroupUser(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

func PostGroupPermission(w http.ResponseWriter, r *http.Request) {
	response := services.CreateGroupPermission(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

func DeleteGroupPermission(w http.ResponseWriter, r *http.Request) {
	response := services.DeleteGroupPermission(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// package admin

// import (
// 	"encoding/json"
// 	"io/ioutil"
// 	"net/http"
// 	"strconv"

// 	"github.com/go-chi/chi"

// 	"github.com/cryo-management/api/common"
// 	"github.com/cryo-management/api/models"
// 	services "github.com/cryo-management/api/services/admin"

// 	"github.com/go-chi/render"
// )

// func PostGroup(w http.ResponseWriter, r *http.Request) {
// 	group := new(models.Group)
// 	body, err := ioutil.ReadAll(r.Body)
// 	err = json.Unmarshal(body, &group)
// 	if err != nil {
// 		render.Status(r, http.StatusInternalServerError)
// 		render.JSON(w, r, common.NewResponseError(common.ErrorParsingRequest, "PostGroup unmarshal body", err.Error()))
// 		return
// 	}

// 	groupService := new(services.GroupService)
// 	err = groupService.Create(group)
// 	if err != nil {
// 		render.Status(r, http.StatusInternalServerError)
// 		render.JSON(w, r, common.NewResponseError(common.ErrorInsertingRecord, "PostGroup creating", err.Error()))
// 		return
// 	}

// 	render.JSON(w, r, group)
// }

// func GetGroup(w http.ResponseWriter, r *http.Request) {
// 	group := new(models.Group)
// 	id := string(chi.URLParam(r, "group_id"))

// 	groupService := new(services.GroupService)
// 	err := groupService.Load(group, id)
// 	if err != nil {
// 		render.Status(r, http.StatusInternalServerError)
// 		render.JSON(w, r, common.NewResponseError(common.ErrorReturningData, "GetGroup load", err.Error()))
// 		return
// 	}

// 	render.JSON(w, r, group)
// }

// func GetAllGroups(w http.ResponseWriter, r *http.Request) {
// 	groups := new(models.Groups)

// 	groupService := new(services.GroupService)
// 	err := groupService.LoadAll(groups)
// 	if err != nil {
// 		render.Status(r, http.StatusInternalServerError)
// 		render.JSON(w, r, common.NewResponseError(common.ErrorReturningData, "GetAllGroups load", err.Error()))
// 		return
// 	}

// 	render.JSON(w, r, groups)
// }

// func DeleteGroup(w http.ResponseWriter, r *http.Request) {
// 	group := new(models.Group)
// 	id := string(chi.URLParam(r, "group_id"))

// 	groupService := new(services.GroupService)
// 	err := groupService.Delete(group, id)
// 	if err != nil {
// 		render.Status(r, http.StatusInternalServerError)
// 		render.JSON(w, r, common.NewResponseError(common.ErrorDeletingData, "DeleteGroup delete group", err.Error()))
// 		return
// 	}

// 	render.JSON(w, r, id)
// }

// func AddGroupUser(w http.ResponseWriter, r *http.Request) {
// 	groupUser := new(models.GroupUser)
// 	body, err := ioutil.ReadAll(r.Body)
// 	err = json.Unmarshal(body, &groupUser)
// 	if err != nil {
// 		render.Status(r, http.StatusInternalServerError)
// 		render.JSON(w, r, common.NewResponseError(common.ErrorParsingRequest, "AddUser unmarshal body", err.Error()))
// 		return
// 	}

// 	groupService := new(services.GroupService)
// 	err = groupService.CreateGroupUser(groupUser)
// 	if err != nil {
// 		render.Status(r, http.StatusInternalServerError)
// 		render.JSON(w, r, common.NewResponseError(common.ErrorInsertingRecord, "AddUser creating", err.Error()))
// 		return
// 	}

// 	render.JSON(w, r, groupUser)
// }

// func RemoveGroupUser(w http.ResponseWriter, r *http.Request) {
// 	groupUser := new(models.GroupUser)
// 	groupID := string(chi.URLParam(r, "group_id"))
// 	userID := string(chi.URLParam(r, "user_id"))

// 	groupUser.GroupID = groupID
// 	groupUser.UserID = userID

// 	groupService := new(services.GroupService)
// 	err := groupService.DeleteGoupUser(groupUser)
// 	if err != nil {
// 		render.Status(r, http.StatusInternalServerError)
// 		render.JSON(w, r, common.NewResponseError(common.ErrorDeletingData, "RemoveUser delete group user", err.Error()))
// 		return
// 	}

// 	render.JSON(w, r, groupUser)
// }

// func AddPermission(w http.ResponseWriter, r *http.Request) {
// 	groupPermission := new(models.GroupPermission)
// 	body, err := ioutil.ReadAll(r.Body)
// 	err = json.Unmarshal(body, &groupPermission)
// 	if err != nil {
// 		render.Status(r, http.StatusInternalServerError)
// 		render.JSON(w, r, common.NewResponseError(common.ErrorParsingRequest, "AddPermission unmarshal body", err.Error()))
// 		return
// 	}

// 	groupService := new(services.GroupService)
// 	err = groupService.CreatePermission(groupPermission)
// 	if err != nil {
// 		render.Status(r, http.StatusInternalServerError)
// 		render.JSON(w, r, common.NewResponseError(common.ErrorInsertingRecord, "AddPermission creating", err.Error()))
// 		return
// 	}

// 	render.JSON(w, r, groupPermission)
// }

// func RemovePermission(w http.ResponseWriter, r *http.Request) {
// 	groupPermission := new(models.GroupPermission)
// 	groupID := string(chi.URLParam(r, "group_id"))
// 	structureID := string(chi.URLParam(r, "structure_id"))
// 	permissionType, err := strconv.Atoi(chi.URLParam(r, "type"))
// 	if err != nil {
// 		render.Status(r, http.StatusInternalServerError)
// 		render.JSON(w, r, common.NewResponseError(common.ErrorDeletingData, "GroupPermission converting string to int", err.Error()))
// 		return
// 	}

// 	groupPermission.GroupID = groupID
// 	groupPermission.StructureID = structureID
// 	groupPermission.Type = permissionType

// 	groupService := new(services.GroupService)
// 	err = groupService.DeletePermission(groupPermission)
// 	if err != nil {
// 		render.Status(r, http.StatusInternalServerError)
// 		render.JSON(w, r, common.NewResponseError(common.ErrorDeletingData, "GroupPermission delete group user", err.Error()))
// 		return
// 	}

// 	render.JSON(w, r, groupPermission)
// }
