package services

import (
	"net/http"

	"github.com/andreluzz/go-sql-builder/builder"
	"github.com/go-chi/chi"

	"github.com/cryo-management/api/models"
)

// CreateLookup persists the request body creating a new object in the database
func CreateLookup(r *http.Request) *Response {
	lookup := models.Lookup{}

	return create(r, &lookup, "CreateLookup", models.TableLookups)
}

// LoadAllLookups return all instances from the object
func LoadAllLookups(r *http.Request) *Response {
	lookups := []models.Lookup{}

	return load(r, &lookups, "LoadAllLookups", models.TableLookups, nil)
}

// LoadLookup return only one object from the database
func LoadLookup(r *http.Request) *Response {
	lookup := models.Lookup{}
	lookupID := chi.URLParam(r, "lookup_id")
	condition := builder.Equal("lookups.id", lookupID)

	return load(r, &lookup, "LoadLookup", models.TableLookups, condition)
}

// UpdateLookup updates object data in the database
func UpdateLookup(r *http.Request) *Response {
	lookupID := chi.URLParam(r, "lookup_id")
	condition := builder.Equal("lookups.id", lookupID)
	lookup := models.Lookup{
		ID: lookupID,
	}

	return update(r, &lookup, "UpdateLookup", models.TableLookups, condition)
}

// DeleteLookup deletes object from the database
func DeleteLookup(r *http.Request) *Response {
	lookupID := chi.URLParam(r, "lookup_id")
	condition := builder.Equal("lookups.id", lookupID)

	return remove(r, "DeleteLookup", models.TableLookups, condition)
}
