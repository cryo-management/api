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
	err := db.LoadStruct(models.TableGroups, &groups, nil)
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
	err := db.LoadStruct(models.TableGroups, group, builder.Equal("groups.id", groupID))
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

//DeleteGroup deletes object from the database
func DeleteGroup(r *http.Request) *Response {
	return nil
}
