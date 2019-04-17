package services

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/andreluzz/go-sql-builder/builder"
	"github.com/andreluzz/go-sql-builder/db"
	"github.com/go-chi/chi"

	"github.com/cryo-management/api/models"
)

//CreateGroup persists the request body creating a new group in the database
func CreateGroup(r *http.Request) *Response {
	response := NewResponse()
	body, _ := ioutil.ReadAll(r.Body)
	group := &models.Group{}
	err := json.Unmarshal(body, group)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorParsingRequest, "CreateGroup unmarshal body", err.Error()))
		return response
	}

	id, err := db.InsertStruct(models.TableGroups, group)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorInsertingRecord, "CreateGroup create", err.Error()))
		return response
	}
	group.ID = id

	//TODO change language_code get from request
	err = models.CreateTranslationsFromStruct(models.TableGroups, r.Header.Get("languageCode"), group)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorInsertingRecord, "CreateGroup create translation", err.Error()))
		return response
	}

	response.Data = group

	return response
}

func LoadAllGroups(r *http.Request) *Response {
	response := NewResponse()

	groups := []models.Group{}
	jsonBytes, err := db.LoadStruct(models.TableGroups, groups, nil)
	json.Unmarshal(jsonBytes, &groups)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorLoadingData, "LoadAllGroups", err.Error()))
		return response
	}
	response.Data = groups
	return response
}

func LoadGroup(r *http.Request) *Response {
	response := NewResponse()
	groupID := chi.URLParam(r, "group_id")
	group := &models.Group{}
	jsonBytes, err := db.LoadStruct(models.TableGroups, group, builder.Equal("groups.id", groupID))
	json.Unmarshal(jsonBytes, group)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorLoadingData, "LoadGroup", err.Error()))
		return response
	}
	response.Data = group
	return response
}

func UpdateGroup(r *http.Request) *Response {
	return nil
}

func DeleteGroup(r *http.Request) *Response {
	return nil
}

//CreateGroupUser persists the request body creating a new groupUser in the database
func CreateGroupUser(r *http.Request) *Response {
	response := NewResponse()
	body, _ := ioutil.ReadAll(r.Body)
	groupUser := &models.GroupUser{}
	err := json.Unmarshal(body, groupUser)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorParsingRequest, "CreateGroupUser unmarshal body", err.Error()))
		return response
	}

	id, err := db.InsertStruct(models.TableGroupsUsers, groupUser)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorInsertingRecord, "CreateGroupUser create", err.Error()))
		return response
	}
	groupUser.ID = id

	response.Data = groupUser

	return response
}

func DeleteGroupUser(r *http.Request) *Response {
	return nil
}

//CreateGroupPermission persists the request body creating a new groupPermission in the database
func CreateGroupPermission(r *http.Request) *Response {
	response := NewResponse()
	body, _ := ioutil.ReadAll(r.Body)
	groupPermission := &models.GroupPermission{}
	err := json.Unmarshal(body, groupPermission)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorParsingRequest, "CreateGroupPermission unmarshal body", err.Error()))
		return response
	}

	id, err := db.InsertStruct(models.TableGroupsPermissions, groupPermission)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorInsertingRecord, "CreateGroupPermission create", err.Error()))
		return response
	}
	groupPermission.ID = id

	response.Data = groupPermission

	return response
}

func DeleteGroupPermission(r *http.Request) *Response {
	return nil
}

// package services

// import (
// 	"github.com/cryo-management/api/models"
// )

// type GroupService struct{}

// func (g *GroupService) Create(group *models.Group) error {
// 	err := group.Create()
// 	if err != nil {
// 		return err
// 	}

// 	// translationService := new(TranslationService)
// 	// err = translationService.Create(*group, group.ID)
// 	// if err != nil {
// 	// 	return err
// 	// }

// 	return nil
// }

// func (g *GroupService) Load(group *models.Group, id string) error {
// 	err := group.Load(id)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (g *GroupService) LoadAll(groups *models.Groups) error {
// 	err := groups.Load()
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (g *GroupService) Delete(group *models.Group, id string) error {
// 	err := group.Delete(id)
// 	if err != nil {
// 		return err
// 	}

// 	// translationService := new(TranslationService)
// 	// err = translationService.DeleteByStructureID(id)
// 	// if err != nil {
// 	// 	return err
// 	// }

// 	groupUser := new(models.GroupUser)
// 	groupUser.GroupID = id

// 	groupService := new(GroupService)
// 	err = groupService.DeleteGoupUser(groupUser)
// 	if err != nil {
// 		return err
// 	}

// 	groupPermission := new(models.GroupPermission)
// 	groupPermission.GroupID = id

// 	err = groupService.DeletePermission(groupPermission)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (g *GroupService) CreateGroupUser(groupUser *models.GroupUser) error {
// 	err := groupUser.Create()
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (g *GroupService) DeleteGoupUser(groupUser *models.GroupUser) error {
// 	err := groupUser.Delete()
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (g *GroupService) CreatePermission(groupPermission *models.GroupPermission) error {
// 	err := groupPermission.Create()
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (g *GroupService) DeletePermission(groupPermission *models.GroupPermission) error {
// 	err := groupPermission.Delete()
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
