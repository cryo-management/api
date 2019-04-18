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

//CreateField persists the request body creating a new object in the database
func CreateField(r *http.Request) *Response {
	field := models.Field{}

	return create(r, &field, "CreateField", models.TableFields)
}

//LoadAllFields return all instances from the object
func LoadAllFields(r *http.Request) *Response {
	fields := []models.Field{}
	schemaID := chi.URLParam(r, "schema_id")
	condition := builder.Equal("fields.schema_id", schemaID)

	return load(r, &fields, "LoadAllFields", models.TableFields, condition)
}

//LoadField return only one object from the database
func LoadField(r *http.Request) *Response {
	field := models.Field{}
	fieldID := chi.URLParam(r, "field_id")
	condition := builder.Equal("fields.id", fieldID)

	return load(r, &field, "LoadAField", models.TableFields, condition)
}

//UpdateField updates object data in the database
func UpdateField(r *http.Request) *Response {
	response := NewResponse()
	fieldID := chi.URLParam(r, "field_id")
	field := &models.Field{}
	body, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(body, field)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorParsingRequest, "UpdateField unmarshal body", err.Error()))

		return response
	}

	condition := builder.Equal("fields.id", fieldID)
	columns := getColumnsFromBody(body)

	err = db.UpdateStruct(models.TableFields, field, condition, columns...)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorInsertingRecord, "UpdateField", err.Error()))

		return response
	}

	return response
}

//DeleteField deletes object from the database
func DeleteField(r *http.Request) *Response {
	fieldID := chi.URLParam(r, "field_id")
	condition := builder.Equal("fields.id", fieldID)

	return delete(r, "DeleteField", models.TableFields, condition)
}
