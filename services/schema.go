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

//CreateSchema persists the request body creating a new object in the database
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

	//Tarefas
	//	1. criar campos padrão
	//		Ex.: createTask(fk_id, fk_table, sequence, status, "Criação de campos padrão", api_method, api_url, json_data)
	//		Worker --> api_method: post | api_url: /api/v1/admin/schemas/{schema_id}/fields | body: json_data
	//	2. Criar view padrão
	//		Ex.: createTask(fk_id, fk_table, sequence, status, "Criação de views", api_method, api_url, json_data)
	//		Worker --> api_method: post | api_url: /api/v1/admin/schemas/{schema_id}/views | body: json_data

	return response
}

//LoadAllSchemas return all instances from the object
func LoadAllSchemas(r *http.Request) *Response {
	schemas := []models.Schema{}
	return load(r, &schemas, "LoadAllSchemas", models.TableSchemas, nil)
}

//LoadSchema return only one object from the database
func LoadSchema(r *http.Request) *Response {
	schema := models.Schema{}
	schemaID := chi.URLParam(r, "schema_id")
	return load(r, &schema, "LoadASchema", models.TableSchemas, builder.Equal("schemas.id", schemaID))
}

//UpdateSchema updates object data in the database
func UpdateSchema(r *http.Request) *Response {
	response := NewResponse()
	schemaID := chi.URLParam(r, "schema_id")
	schema := &models.Schema{}
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, schema)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorParsingRequest, "UpdateSchema unmarshal body", err.Error()))
		return response
	}

	// TODO: Extract this to a method getBodyFields(body) []string
	jsonMap := make(map[string]interface{})
	json.Unmarshal(body, &jsonMap)
	fields := []string{}
	for k := range jsonMap {
		fields = append(fields, k)
	}

	err = db.UpdateStruct(models.TableSchemas, schema, builder.Equal("schemas.id", schemaID), fields...)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorInsertingRecord, "UpdateSchema", err.Error()))
		return response
	}

	return response
}

//DeleteSchema deletes object from the database
func DeleteSchema(r *http.Request) *Response {
	response := NewResponse()
	schemaID := chi.URLParam(r, "schema_id")
	err := db.DeleteStruct(models.TableSchemas, builder.Equal("schemas.id", schemaID))
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorInsertingRecord, "UpdateSchema", err.Error()))
		return response
	}
	return response
}
