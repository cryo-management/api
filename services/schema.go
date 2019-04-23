package services

import (
	"net/http"

	"github.com/andreluzz/go-sql-builder/builder"
	"github.com/go-chi/chi"

	"github.com/cryo-management/api/models"
)

// CreateSchema persists the request body creating a new object in the database
func CreateSchema(r *http.Request) *Response {
	schema := models.Schema{}

	return create(r, &schema, "CreateSchema", models.TableCoreSchemas)
}

// LoadAllSchemas return all instances from the object
func LoadAllSchemas(r *http.Request) *Response {
	schemas := []models.Schema{}

	return load(r, &schemas, "LoadAllSchemas", models.TableCoreSchemas, nil)
}

// LoadSchema return only one object from the database
func LoadSchema(r *http.Request) *Response {
	schema := models.Schema{}
	schemaID := chi.URLParam(r, "schema_id")
	condition := builder.Equal("schemas.id", schemaID)

	return load(r, &schema, "LoadSchema", models.TableCoreSchemas, condition)
}

// UpdateSchema updates object data in the database
func UpdateSchema(r *http.Request) *Response {
	schemaID := chi.URLParam(r, "schema_id")
	condition := builder.Equal("schemas.id", schemaID)
	schema := models.Schema{
		ID: schemaID,
	}

	return update(r, &schema, "UpdateSchema", models.TableCoreSchemas, condition)
}

// DeleteSchema deletes object from the database
func DeleteSchema(r *http.Request) *Response {
	schemaID := chi.URLParam(r, "schema_id")
	condition := builder.Equal("schemas.id", schemaID)

	return remove(r, "DeleteSchema", models.TableCoreSchemas, condition)
}
