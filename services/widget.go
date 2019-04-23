package services

import (
	"net/http"

	"github.com/andreluzz/go-sql-builder/builder"
	"github.com/go-chi/chi"

	"github.com/cryo-management/api/models"
)

// CreateWidget persists the request body creating a new object in the database
func CreateWidget(r *http.Request) *Response {
	widget := models.Widget{}

	return create(r, &widget, "CreateWidget", models.TableCoreWidgets)
}

// LoadAllWidgets return all instances from the object
func LoadAllWidgets(r *http.Request) *Response {
	widgets := []models.Widget{}

	return load(r, &widgets, "LoadAllWidgets", models.TableCoreWidgets, nil)
}

// LoadWidget return only one object from the database
func LoadWidget(r *http.Request) *Response {
	widget := models.Widget{}
	widgetID := chi.URLParam(r, "widget_id")
	condition := builder.Equal("widgets.id", widgetID)

	return load(r, &widget, "LoadWidget", models.TableCoreWidgets, condition)
}

// UpdateWidget updates object data in the database
func UpdateWidget(r *http.Request) *Response {
	widgetID := chi.URLParam(r, "widget_id")
	condition := builder.Equal("widgets.id", widgetID)
	widget := models.Widget{
		ID: widgetID,
	}

	return update(r, &widget, "UpdateWidget", models.TableCoreWidgets, condition)
}

// DeleteWidget deletes object from the database
func DeleteWidget(r *http.Request) *Response {
	widgetID := chi.URLParam(r, "widget_id")
	condition := builder.Equal("widgets.id", widgetID)

	return remove(r, "DeleteWidget", models.TableCoreWidgets, condition)
}
