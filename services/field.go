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

//CreateField persists the request body creating a new field in the database
func CreateField(r *http.Request) *Response {
	response := NewResponse()
	body, _ := ioutil.ReadAll(r.Body)
	field := &models.Field{}
	err := json.Unmarshal(body, field)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorParsingRequest, "CreateField unmarshal body", err.Error()))
		return response
	}

	id, err := db.InsertStruct(models.TableFields, field)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorInsertingRecord, "CreateField create", err.Error()))
		return response
	}
	field.ID = id

	err = models.CreateTranslationsFromStruct(models.TableFields, r.Header.Get("languageCode"), field)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorInsertingRecord, "CreateField create translation", err.Error()))
		return response
	}

	response.Data = field

	return response
}

func LoadAllFields(r *http.Request) *Response {
	response := NewResponse()
	schemaID := chi.URLParam(r, "schema_id")
	fields := []models.Field{}
	err := db.LoadStruct(models.TableFields, &fields, builder.Equal("fields.schema_id", schemaID))
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorLoadingData, "LoadAllFields loaging data", err.Error()))
		return response
	}
	response.Data = fields
	return response
}

func LoadField(r *http.Request) *Response {
	response := NewResponse()
	fieldID := chi.URLParam(r, "field_id")
	field := &models.Field{}
	err := db.LoadStruct(models.TableFields, field, builder.Equal("fields.id", fieldID))
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorLoadingData, "GetField", err.Error()))
		return response
	}
	response.Data = field
	return response
}

func UpdateField(r *http.Request) *Response {
	return nil
}

//DeleteField deletes object from the database
func DeleteField(r *http.Request) *Response {
	return nil
}
