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

//CreateLookup persists the request body creating a new object in the database
func CreateLookup(r *http.Request) *Response {
	lookup := models.Lookup{}

	return create(r, &lookup, "CreateLookup", models.TableLookups)
}

//LoadAllLookups return all instances from the object
func LoadAllLookups(r *http.Request) *Response {
	lookups := []models.Lookup{}

	return load(r, &lookups, "LoadAllLookups", models.TableLookups, nil)
}

//LoadLookup return only one object from the database
func LoadLookup(r *http.Request) *Response {
	lookup := models.Lookup{}
	lookupID := chi.URLParam(r, "lookup_id")
	condition := builder.Equal("lookups.id", lookupID)

	return load(r, &lookup, "LoadLookup", models.TableLookups, condition)
}

//UpdateLookup updates object data in the database
func UpdateLookup(r *http.Request) *Response {
	response := NewResponse()
	lookupID := chi.URLParam(r, "lookup_id")
	lookup := &models.Lookup{}
	body, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(body, lookup)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorParsingRequest, "UpdateLookup unmarshal body", err.Error()))

		return response
	}

	condition := builder.Equal("lookups.id", lookupID)
	columns := getColumnsFromBody(body)

	err = db.UpdateStruct(models.TableLookups, lookup, condition, columns...)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorInsertingRecord, "UpdateLookup", err.Error()))

		return response
	}

	return response
}

//DeleteLookup deletes object from the database
func DeleteLookup(r *http.Request) *Response {
	lookupID := chi.URLParam(r, "lookup_id")
	condition := builder.Equal("lookups.id", lookupID)

	return remove(r, "DeleteLookup", models.TableLookups, condition)
}
