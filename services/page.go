package services

import (
	"net/http"

	"github.com/andreluzz/go-sql-builder/builder"
	"github.com/go-chi/chi"

	"github.com/cryo-management/api/models"
)

// CreatePage persists the request body creating a new object in the database
func CreatePage(r *http.Request) *Response {
	page := models.Page{}

	return create(r, &page, "CreatePage", models.TableCoreSchPages)
}

// LoadAllPages return all instances from the object
func LoadAllPages(r *http.Request) *Response {
	pages := []models.Page{}
	viewID := chi.URLParam(r, "view_id")
	condition := builder.Equal("pages.view_id", viewID)

	return load(r, &pages, "LoadAllPages", models.TableCoreSchPages, condition)
}

// LoadPage return only one object from the database
func LoadPage(r *http.Request) *Response {
	page := models.Page{}
	pageID := chi.URLParam(r, "page_id")
	condition := builder.Equal("pages.id", pageID)

	return load(r, &page, "LoadPage", models.TableCoreSchPages, condition)
}

// UpdatePage updates object data in the database
func UpdatePage(r *http.Request) *Response {
	pageID := chi.URLParam(r, "page_id")
	condition := builder.Equal("pages.id", pageID)
	page := models.Page{
		ID: pageID,
	}

	return update(r, &page, "UpdatePage", models.TableCoreSchPages, condition)
}

// DeletePage deletes object from the database
func DeletePage(r *http.Request) *Response {
	pageID := chi.URLParam(r, "page_id")
	condition := builder.Equal("pages.id", pageID)

	return remove(r, "DeletePage", models.TableCoreSchPages, condition)
}
