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

//CreateSchema persists the request body creating a new schema in the database
func CreateSchema(r *http.Request) *Response {
	response := NewResponse()
	body, _ := ioutil.ReadAll(r.Body)
	schema := &models.Schema{}
	err := json.Unmarshal(body, schema)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorParsingRequest, "CreateSchema unmarshal body", err.Error()))
		return response
	}

	id, err := db.InsertStruct(models.TableSchemas, schema)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorInsertingRecord, "CreateSchema create", err.Error()))
		return response
	}
	schema.ID = id

	//TODO change language_code get from request
	err = models.CreateTranslationsFromStruct(models.TableSchemas, r.Header.Get("languageCode"), schema)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorInsertingRecord, "CreateSchema create translation", err.Error()))
		return response
	}

	response.Data = schema

	return response
}

func LoadAllSchemas(r *http.Request) *Response {
	response := NewResponse()

	schemas := []models.Schema{}
	jsonBytes, err := db.LoadStruct(models.TableSchemas, schemas, nil)
	json.Unmarshal(jsonBytes, &schemas)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorLoadingData, "LoadAllSchemas", err.Error()))
		return response
	}
	response.Data = schemas
	return response
}

func LoadSchema(r *http.Request) *Response {
	response := NewResponse()
	schemaID := chi.URLParam(r, "schema_id")
	schema := &models.Schema{}
	jsonBytes, err := db.LoadStruct(models.TableSchemas, schema, builder.Equal("schemas.id", schemaID))
	json.Unmarshal(jsonBytes, schema)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorLoadingData, "LoadSchema", err.Error()))
		return response
	}
	response.Data = schema
	return response
}

func UpdateSchema(r *http.Request) *Response {
	return nil
}

func DeleteSchema(r *http.Request) *Response {
	return nil
}
