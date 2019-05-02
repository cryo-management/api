package resources

import (
	"fmt"
	"net/http"

	"github.com/andreluzz/go-sql-builder/builder"
	"github.com/go-chi/chi"

	"github.com/cryo-management/api/models"
	"github.com/cryo-management/api/services"
)

// CreateLookup persists the request body creating a new object in the database
func CreateLookup(r *http.Request) *services.Response {
	lookup := models.Lookup{}

	return services.Create(r, &lookup, "CreateLookup", models.TableCoreLookups)
}

// LoadAllLookups return all instances from the object
func LoadAllLookups(r *http.Request) *services.Response {
	lookups := []models.Lookup{}

	return services.Load(r, &lookups, "LoadAllLookups", models.TableCoreLookups, nil)
}

// LoadLookup return only one object from the database
func LoadLookup(r *http.Request) *services.Response {
	lookup := models.Lookup{}
	lookupID := chi.URLParam(r, "lookup_id")
	lookupIDColumn := fmt.Sprintf("%s.id", models.TableCoreLookups)
	condition := builder.Equal(lookupIDColumn, lookupID)

	return services.Load(r, &lookup, "LoadLookup", models.TableCoreLookups, condition)
}

// UpdateLookup updates object data in the database
func UpdateLookup(r *http.Request) *services.Response {
	lookupID := chi.URLParam(r, "lookup_id")
	lookupIDColumn := fmt.Sprintf("%s.id", models.TableCoreLookups)
	condition := builder.Equal(lookupIDColumn, lookupID)
	lookup := models.Lookup{
		ID: lookupID,
	}

	return services.Update(r, &lookup, "UpdateLookup", models.TableCoreLookups, condition)
}

// DeleteLookup deletes object from the database
func DeleteLookup(r *http.Request) *services.Response {
	lookupID := chi.URLParam(r, "lookup_id")
	lookupIDColumn := fmt.Sprintf("%s.id", models.TableCoreLookups)
	condition := builder.Equal(lookupIDColumn, lookupID)

	return services.Remove(r, "DeleteLookup", models.TableCoreLookups, condition)
}

// CreateLookupOption persists the request body creating a new object in the database
func CreateLookupOption(r *http.Request) *services.Response {
	lookupOption := models.LookupOption{}

	return services.Create(r, &lookupOption, "CreateLookupOption", models.TableCoreLkpOptions)
}

// LoadAllLookupOptions return all instances from the object
func LoadAllLookupOptions(r *http.Request) *services.Response {
	lookupOptions := []models.LookupOption{}
	lookupID := chi.URLParam(r, "lookup_id")
	lookupIDColumn := fmt.Sprintf("%s.lookup_id", models.TableCoreLkpOptions)
	condition := builder.Equal(lookupIDColumn, lookupID)

	return services.Load(r, &lookupOptions, "LoadAllLookupOptions", models.TableCoreLkpOptions, condition)
}

// LoadLookupOption return only one object from the database
func LoadLookupOption(r *http.Request) *services.Response {
	lookupOption := models.LookupOption{}
	lookupOptionID := chi.URLParam(r, "lookup_option_id")
	lookupOptionIDColumn := fmt.Sprintf("%s.id", models.TableCoreLkpOptions)
	condition := builder.Equal(lookupOptionIDColumn, lookupOptionID)

	return services.Load(r, &lookupOption, "LoadLookupOption", models.TableCoreLkpOptions, condition)
}

// UpdateLookupOption updates object data in the database
func UpdateLookupOption(r *http.Request) *services.Response {
	lookupOptionID := chi.URLParam(r, "lookup_option_id")
	lookupOptionIDColumn := fmt.Sprintf("%s.id", models.TableCoreLkpOptions)
	condition := builder.Equal(lookupOptionIDColumn, lookupOptionID)
	lookupOption := models.LookupOption{
		ID: lookupOptionID,
	}

	return services.Update(r, &lookupOption, "UpdateLookupOption", models.TableCoreLkpOptions, condition)
}

// DeleteLookupOption deletes object from the database
func DeleteLookupOption(r *http.Request) *services.Response {
	lookupOptionID := chi.URLParam(r, "lookup_option_id")
	lookupOptionIDColumn := fmt.Sprintf("%s.id", models.TableCoreLkpOptions)
	condition := builder.Equal(lookupOptionIDColumn, lookupOptionID)

	return services.Remove(r, "DeleteLookupOption", models.TableCoreLkpOptions, condition)
}
