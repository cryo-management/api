package resources

import (
	"fmt"
	"net/http"

	"github.com/andreluzz/go-sql-builder/builder"
	"github.com/go-chi/chi"

	"github.com/cryo-management/api/models"
	"github.com/cryo-management/api/services"
)

// CreateWidget persists the request body creating a new object in the database
func CreateWidget(r *http.Request) *services.Response {
	widget := models.Widget{}

	return services.Create(r, &widget, "CreateWidget", models.TableCoreWidgets)
}

// LoadAllWidgets return all instances from the object
func LoadAllWidgets(r *http.Request) *services.Response {
	widgets := []models.Widget{}

	return services.Load(r, &widgets, "LoadAllWidgets", models.TableCoreWidgets, nil)
}

// LoadWidget return only one object from the database
func LoadWidget(r *http.Request) *services.Response {
	widget := models.Widget{}
	widgetID := chi.URLParam(r, "widget_id")
	widgetIDColumn := fmt.Sprintf("%s.id", models.TableCoreWidgets)
	condition := builder.Equal(widgetIDColumn, widgetID)

	return services.Load(r, &widget, "LoadWidget", models.TableCoreWidgets, condition)
}

// UpdateWidget updates object data in the database
func UpdateWidget(r *http.Request) *services.Response {
	widgetID := chi.URLParam(r, "widget_id")
	widgetIDColumn := fmt.Sprintf("%s.id", models.TableCoreWidgets)
	condition := builder.Equal(widgetIDColumn, widgetID)
	widget := models.Widget{
		ID: widgetID,
	}

	return services.Update(r, &widget, "UpdateWidget", models.TableCoreWidgets, condition)
}

// DeleteWidget deletes object from the database
func DeleteWidget(r *http.Request) *services.Response {
	widgetID := chi.URLParam(r, "widget_id")
	widgetIDColumn := fmt.Sprintf("%s.id", models.TableCoreWidgets)
	condition := builder.Equal(widgetIDColumn, widgetID)

	return services.Remove(r, "DeleteWidget", models.TableCoreWidgets, condition)
}
