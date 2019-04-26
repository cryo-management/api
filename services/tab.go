package services

import (
	"fmt"
	"net/http"

	"github.com/andreluzz/go-sql-builder/builder"
	"github.com/go-chi/chi"

	"github.com/cryo-management/api/models"
)

// CreateTab persists the request body creating a new object in the database
func CreateTab(r *http.Request) *Response {
	tab := models.Tab{}

	return create(r, &tab, "CreateTab", models.TableCoreSchPagSecTabs)
}

// LoadAllTabs return all instances from the object
func LoadAllTabs(r *http.Request) *Response {
	tabs := []models.Tab{}
	sectionID := chi.URLParam(r, "section_id")
	sectionIDColumn := fmt.Sprintf("%s.section_id", models.TableCoreSchPagSecTabs)
	condition := builder.Equal(sectionIDColumn, sectionID)

	return load(r, &tabs, "LoadAllTabs", models.TableCoreSchPagSecTabs, condition)
}

// LoadTab return only one object from the database
func LoadTab(r *http.Request) *Response {
	tab := models.Tab{}
	tabID := chi.URLParam(r, "tab_id")
	tabIDColumn := fmt.Sprintf("%s.id", models.TableCoreSchPagSecTabs)
	condition := builder.Equal(tabIDColumn, tabID)

	return load(r, &tab, "LoadTab", models.TableCoreSchPagSecTabs, condition)
}

// UpdateTab updates object data in the database
func UpdateTab(r *http.Request) *Response {
	tabID := chi.URLParam(r, "tab_id")
	tabIDColumn := fmt.Sprintf("%s.id", models.TableCoreSchPagSecTabs)
	condition := builder.Equal(tabIDColumn, tabID)
	tab := models.Tab{
		ID: tabID,
	}

	return update(r, &tab, "UpdateTab", models.TableCoreSchPagSecTabs, condition)
}

// DeleteTab deletes object from the database
func DeleteTab(r *http.Request) *Response {
	tabID := chi.URLParam(r, "tab_id")
	tabIDColumn := fmt.Sprintf("%s.id", models.TableCoreSchPagSecTabs)
	condition := builder.Equal(tabIDColumn, tabID)

	return remove(r, "DeleteTab", models.TableCoreSchPagSecTabs, condition)
}
