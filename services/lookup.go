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

//CreateLookup persists the request body creating a new lookup in the database
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

func LoadAllLookups(r *http.Request) *Response {
	response := NewResponse()

	lookups := []models.Lookup{}
	jsonBytes, err := db.LoadStruct(models.TableLookups, lookups, nil)
	json.Unmarshal(jsonBytes, &lookups)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorLoadingData, "LoadAllLookups", err.Error()))
		return response
	}
	response.Data = lookups
	return response
}

func LoadLookup(r *http.Request) *Response {
	response := NewResponse()
	lookupID := chi.URLParam(r, "lookup_id")
	lookup := &models.Lookup{}
	jsonBytes, err := db.LoadStruct(models.TableLookups, lookup, builder.Equal("lookups.id", lookupID))
	json.Unmarshal(jsonBytes, lookup)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorLoadingData, "LoadLookup", err.Error()))
		return response
	}
	response.Data = lookup
	return response
}

func UpdateLookup(r *http.Request) *Response {
	return nil
}

func DeleteLookup(r *http.Request) *Response {
	return nil
}
