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
	schema := models.Schema{}

	return create(r, &schema, "CreateSchema", models.TableSchemas)
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
	condition := builder.Equal("schemas.id", schemaID)

	return load(r, &schema, "LoadASchema", models.TableSchemas, condition)
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

	condition := builder.Equal("schemas.id", schemaID)
	columns := getColumnsFromBody(body)

	err = db.UpdateStruct(models.TableSchemas, schema, condition, columns...)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorInsertingRecord, "UpdateSchema", err.Error()))

		return response
	}

	return response
}

//DeleteSchema deletes object from the database
func DeleteSchema(r *http.Request) *Response {
	schemaID := chi.URLParam(r, "schema_id")
	condition := builder.Equal("schemas.id", schemaID)

	return delete(r, "DeleteSchema", models.TableSchemas, condition)
}
