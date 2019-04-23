package services

import (
	"net/http"

	"github.com/andreluzz/go-sql-builder/builder"
	"github.com/go-chi/chi"

	"github.com/cryo-management/api/models"
)

// CreateView persists the request body creating a new object in the database
func CreateView(r *http.Request) *Response {
	view := models.View{}

	return create(r, &view, "CreateView", models.TableCoreSchViews)
}

// LoadAllViews return all instances from the object
func LoadAllViews(r *http.Request) *Response {
	views := []models.View{}
	structureID := chi.URLParam(r, "structure_id")
	condition := builder.Equal("views.structure_id", structureID)

	return load(r, &views, "LoadAllViews", models.TableCoreSchViews, condition)
}

// LoadView return only one object from the database
func LoadView(r *http.Request) *Response {
	view := models.View{}
	viewID := chi.URLParam(r, "view_id")
	condition := builder.Equal("views.id", viewID)

	return load(r, &view, "LoadView", models.TableCoreSchViews, condition)
}

// UpdateView updates object data in the database
func UpdateView(r *http.Request) *Response {
	viewID := chi.URLParam(r, "view_id")
	condition := builder.Equal("views.id", viewID)
	view := models.View{
		ID: viewID,
	}

	return update(r, &view, "UpdateView", models.TableCoreSchViews, condition)
}

// DeleteView deletes object from the database
func DeleteView(r *http.Request) *Response {
	viewID := chi.URLParam(r, "view_id")
	condition := builder.Equal("views.id", viewID)

	return remove(r, "DeleteView", models.TableCoreSchViews, condition)
}
