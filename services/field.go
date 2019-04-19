package services

import (
	"net/http"

	"github.com/andreluzz/go-sql-builder/builder"
	"github.com/go-chi/chi"

	"github.com/cryo-management/api/models"
)

// CreateField persists the request body creating a new object in the database
func CreateField(r *http.Request) *Response {
	field := models.Field{}

	return create(r, &field, "CreateField", models.TableFields)
}

// LoadAllFields return all instances from the object
func LoadAllFields(r *http.Request) *Response {
	fields := []models.Field{}
	schemaID := chi.URLParam(r, "schema_id")
	condition := builder.Equal("fields.schema_id", schemaID)

	return load(r, &fields, "LoadAllFields", models.TableFields, condition)
}

// LoadField return only one object from the database
func LoadField(r *http.Request) *Response {
	field := models.Field{}
	fieldID := chi.URLParam(r, "field_id")
	condition := builder.Equal("fields.id", fieldID)

	return load(r, &field, "LoadField", models.TableFields, condition)
}

// UpdateField updates object data in the database
func UpdateField(r *http.Request) *Response {
	fieldID := chi.URLParam(r, "field_id")
	condition := builder.Equal("fields.id", fieldID)
	field := models.Field{
		ID: fieldID,
	}

	return update(r, &field, "UpdateField", models.TableFields, condition)
}

// DeleteField deletes object from the database
func DeleteField(r *http.Request) *Response {
	fieldID := chi.URLParam(r, "field_id")
	condition := builder.Equal("fields.id", fieldID)

	return remove(r, "DeleteField", models.TableFields, condition)
}
