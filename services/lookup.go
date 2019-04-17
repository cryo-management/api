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
	response := NewResponse()
	body, _ := ioutil.ReadAll(r.Body)
	lookup := &models.Lookup{}
	err := json.Unmarshal(body, lookup)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorParsingRequest, "CreateLookup unmarshal body", err.Error()))
		return response
	}

	id, err := db.InsertStruct(models.TableLookups, lookup)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorInsertingRecord, "CreateLookup create", err.Error()))
		return response
	}
	lookup.ID = id

	//TODO change language_code get from request
	err = models.CreateTranslationsFromStruct(models.TableLookups, r.Header.Get("languageCode"), lookup)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorInsertingRecord, "CreateLookup create translation", err.Error()))
		return response
	}

	response.Data = lookup

	return response
}

//LoadAllLookups return all instances from the object
func LoadAllLookups(r *http.Request) *Response {
	response := NewResponse()

	lookups := []models.Lookup{}
	err := db.LoadStruct(models.TableLookups, &lookups, nil)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorLoadingData, "LoadAllLookups", err.Error()))
		return response
	}
	response.Data = lookups
	return response
}

//LoadLookup return only one object from the database
func LoadLookup(r *http.Request) *Response {
	response := NewResponse()
	lookupID := chi.URLParam(r, "lookup_id")
	lookup := &models.Lookup{}
	err := db.LoadStruct(models.TableLookups, lookup, builder.Equal("lookups.id", lookupID))
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorLoadingData, "LoadLookup", err.Error()))
		return response
	}
	response.Data = lookup
	return response
}

//UpdateLookup updates object data in the database
func UpdateLookup(r *http.Request) *Response {
	return nil
}

//DeleteLookup deletes object from the database
func DeleteLookup(r *http.Request) *Response {
	return nil
}
