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
	response := NewResponse()
	body, _ := ioutil.ReadAll(r.Body)
	lookupOption := &models.LookupOption{}
	err := json.Unmarshal(body, lookupOption)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorParsingRequest, "CreateLookupOption unmarshal body", err.Error()))
		return response
	}

	id, err := db.InsertStruct(models.TableLookupsOptions, lookupOption)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorInsertingRecord, "CreateLookupOption create", err.Error()))
		return response
	}
	lookupOption.ID = id

	err = models.CreateTranslationsFromStruct(models.TableLookupsOptions, r.Header.Get("languageCode"), lookupOption)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorInsertingRecord, "CreateLookupOption create translation", err.Error()))
		return response
	}

	response.Data = lookupOption

	return response
}

//LoadAllLookupOptions return all instances from the object
func LoadAllLookupOptions(r *http.Request) *Response {
	response := NewResponse()
	lookupID := chi.URLParam(r, "lookup_id")
	lookupOptions := []models.LookupOption{}
	err := db.LoadStruct(models.TableLookupsOptions, &lookupOptions, builder.Equal("lookups_options.lookup_id", lookupID))
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorLoadingData, "LoadAllLookupOptions loaging data", err.Error()))
		return response
	}
	response.Data = lookupOptions
	return response
}

//LoadLookupOption return only one object from the database
func LoadLookupOption(r *http.Request) *Response {
	response := NewResponse()
	lookupOptionID := chi.URLParam(r, "lookup_option_id")
	lookupOption := &models.LookupOption{}
	err := db.LoadStruct(models.TableLookupsOptions, lookupOption, builder.Equal("lookups_options.id", lookupOptionID))
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorLoadingData, "GetLookupOption", err.Error()))
		return response
	}
	response.Data = lookupOption
	return response
}

//UpdateLookupOption updates object data in the database
func UpdateLookupOption(r *http.Request) *Response {
	return nil
}

//DeleteLookupOption deletes object from the database
func DeleteLookupOption(r *http.Request) *Response {
	return nil
}
