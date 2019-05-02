package resources

import (
	"fmt"
	"net/http"

	"github.com/andreluzz/go-sql-builder/builder"
	"github.com/go-chi/chi"

	"github.com/cryo-management/api/models"
	"github.com/cryo-management/api/services"
)

// CreatePage persists the request body creating a new object in the database
func CreatePage(r *http.Request) *services.Response {
	page := models.Page{}

	return services.Create(r, &page, "CreatePage", models.TableCoreSchPages)
}

// LoadAllPages return all instances from the object
func LoadAllPages(r *http.Request) *services.Response {
	pages := []models.Page{}
	schemaID := chi.URLParam(r, "schema_id")
	schemaIDColumn := fmt.Sprintf("%s.schema_id", models.TableCoreSchPages)
	condition := builder.Equal(schemaIDColumn, schemaID)

	return services.Load(r, &pages, "LoadAllPages", models.TableCoreSchPages, condition)
}

// LoadPage return only one object from the database
func LoadPage(r *http.Request) *services.Response {
	page := models.Page{}
	pageID := chi.URLParam(r, "page_id")
	pageIDColumn := fmt.Sprintf("%s.id", models.TableCoreSchPages)
	condition := builder.Equal(pageIDColumn, pageID)

	return services.Load(r, &page, "LoadPage", models.TableCoreSchPages, condition)
}

// UpdatePage updates object data in the database
func UpdatePage(r *http.Request) *services.Response {
	pageID := chi.URLParam(r, "page_id")
	pageIDColumn := fmt.Sprintf("%s.id", models.TableCoreSchPages)
	condition := builder.Equal(pageIDColumn, pageID)
	page := models.Page{
		ID: pageID,
	}

	return services.Update(r, &page, "UpdatePage", models.TableCoreSchPages, condition)
}

// DeletePage deletes object from the database
func DeletePage(r *http.Request) *services.Response {
	pageID := chi.URLParam(r, "page_id")
	pageIDColumn := fmt.Sprintf("%s.id", models.TableCoreSchPages)
	condition := builder.Equal(pageIDColumn, pageID)

	return services.Remove(r, "DeletePage", models.TableCoreSchPages, condition)
}
