package services

import (
	"net/http"

	"github.com/andreluzz/go-sql-builder/builder"
	"github.com/go-chi/chi"

	"github.com/cryo-management/api/models"
)

// CreateLookupOption persists the request body creating a new object in the database
func CreateLookupOption(r *http.Request) *Response {
	lookupOption := models.LookupOption{}

	return create(r, &lookupOption, "CreateLookupOption", models.TableLookupsOptions)
}

// LoadAllLookupOptions return all instances from the object
func LoadAllLookupOptions(r *http.Request) *Response {
	lookupOptions := []models.LookupOption{}
	lookupID := chi.URLParam(r, "lookup_id")
	condition := builder.Equal("lookups_options.lookup_id", lookupID)

	return load(r, &lookupOptions, "LoadAllLookupOptions", models.TableLookupsOptions, condition)
}

// LoadLookupOption return only one object from the database
func LoadLookupOption(r *http.Request) *Response {
	lookupOption := models.LookupOption{}
	lookupOptionID := chi.URLParam(r, "lookup_option_id")
	condition := builder.Equal("lookups_options.id", lookupOptionID)

	return load(r, &lookupOption, "LoadLookupOption", models.TableLookupsOptions, condition)
}

// UpdateLookupOption updates object data in the database
func UpdateLookupOption(r *http.Request) *Response {
	lookupOptionID := chi.URLParam(r, "lookup_option_id")
	condition := builder.Equal("lookups_options.id", lookupOptionID)
	lookupOption := models.LookupOption{
		ID: lookupOptionID,
	}

	return update(r, &lookupOption, "UpdateLookupOption", models.TableLookupsOptions, condition)
}

// DeleteLookupOption deletes object from the database
func DeleteLookupOption(r *http.Request) *Response {
	lookupOptionID := chi.URLParam(r, "lookup_option_id")
	condition := builder.Equal("lookups_options.id", lookupOptionID)

	return remove(r, "DeleteLookupOption", models.TableLookupsOptions, condition)
}
