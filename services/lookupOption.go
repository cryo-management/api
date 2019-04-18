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

//CreateLookupOption persists the request body creating a new object in the database
func CreateLookupOption(r *http.Request) *Response {
	lookupOption := models.LookupOption{}

	return create(r, &lookupOption, "CreateLookupOption", models.TableLookupsOptions)
}

//LoadAllLookupOptions return all instances from the object
func LoadAllLookupOptions(r *http.Request) *Response {
	lookupOptions := []models.LookupOption{}
	lookupID := chi.URLParam(r, "lookup_id")
	condition := builder.Equal("lookups_options.lookup_id", lookupID)

	return load(r, &lookupOptions, "LoadAllLookupOptions", models.TableLookupsOptions, condition)
}

//LoadLookupOption return only one object from the database
func LoadLookupOption(r *http.Request) *Response {
	lookupOption := models.LookupOption{}
	lookupOptionID := chi.URLParam(r, "lookup_option_id")
	condition := builder.Equal("lookups_options.id", lookupOptionID)

	return load(r, &lookupOption, "LoadALookupOption", models.TableLookupsOptions, condition)
}

//UpdateLookupOption updates object data in the database
func UpdateLookupOption(r *http.Request) *Response {
	response := NewResponse()
	lookupOptionID := chi.URLParam(r, "lookup_option_id")
	lookupOption := &models.LookupOption{}
	body, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(body, lookupOption)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorParsingRequest, "UpdateLookupOption unmarshal body", err.Error()))

		return response
	}

	condition := builder.Equal("lookups_options.id", lookupOptionID)
	columns := getColumnsFromBody(body)

	err = db.UpdateStruct(models.TableLookupsOptions, lookupOption, condition, columns...)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorInsertingRecord, "UpdateLookupOption", err.Error()))

		return response
	}

	return response
}

//DeleteLookupOption deletes object from the database
func DeleteLookupOption(r *http.Request) *Response {
	lookupOptionID := chi.URLParam(r, "lookup_option_id")
	condition := builder.Equal("lookups_options.id", lookupOptionID)

	return delete(r, "DeleteLookupOption", models.TableLookupsOptions, condition)
}
