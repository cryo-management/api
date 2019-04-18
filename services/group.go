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

//CreateGroup persists the request body creating a new object in the database
func CreateGroup(r *http.Request) *Response {
	group := models.Group{}

	return create(r, &group, "CreateGroup", models.TableGroups)
}

//LoadAllGroups return all instances from the object
func LoadAllGroups(r *http.Request) *Response {
	groups := []models.Group{}

	return load(r, &groups, "LoadAllGroups", models.TableGroups, nil)
}

//LoadGroup return only one object from the database
func LoadGroup(r *http.Request) *Response {
	group := models.Group{}
	groupID := chi.URLParam(r, "group_id")
	condition := builder.Equal("groups.id", groupID)

	return load(r, &group, "LoadAGroup", models.TableGroups, condition)
}

//UpdateGroup updates object data in the database
func UpdateGroup(r *http.Request) *Response {
	response := NewResponse()
	groupID := chi.URLParam(r, "group_id")
	group := &models.Group{}
	body, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(body, group)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorParsingRequest, "UpdateGroup unmarshal body", err.Error()))

		return response
	}

	condition := builder.Equal("groups.id", groupID)
	columns := getColumnsFromBody(body)

	err = db.UpdateStruct(models.TableGroups, group, condition, columns...)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorInsertingRecord, "UpdateGroup", err.Error()))

		return response
	}

	return response
}

//DeleteGroup deletes object from the database
func DeleteGroup(r *http.Request) *Response {
	groupID := chi.URLParam(r, "group_id")
	condition := builder.Equal("groups.id", groupID)

	return delete(r, "DeleteGroup", models.TableGroups, condition)
}
