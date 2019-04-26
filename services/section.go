package services

import (
	"fmt"
	"net/http"

	"github.com/andreluzz/go-sql-builder/builder"
	"github.com/go-chi/chi"

	"github.com/cryo-management/api/models"
)

// CreateSection persists the request body creating a new object in the database
func CreateSection(r *http.Request) *Response {
	section := models.Section{}

	return create(r, &section, "CreateSection", models.TableCoreSchPagSections)
}

// LoadAllSections return all instances from the object
func LoadAllSections(r *http.Request) *Response {
	sections := []models.Section{}
	pageID := chi.URLParam(r, "page_id")
	pageIDColumn := fmt.Sprintf("%s.page_id", models.TableCoreSchPagSections)
	condition := builder.Equal(pageIDColumn, pageID)

	return load(r, &sections, "LoadAllSections", models.TableCoreSchPagSections, condition)
}

// LoadSection return only one object from the database
func LoadSection(r *http.Request) *Response {
	section := models.Section{}
	sectionID := chi.URLParam(r, "section_id")
	sectionIDColumn := fmt.Sprintf("%s.id", models.TableCoreSchPagSections)
	condition := builder.Equal(sectionIDColumn, sectionID)

	return load(r, &section, "LoadSection", models.TableCoreSchPagSections, condition)
}

// UpdateSection updates object data in the database
func UpdateSection(r *http.Request) *Response {
	sectionID := chi.URLParam(r, "section_id")
	sectionIDColumn := fmt.Sprintf("%s.id", models.TableCoreSchPagSections)
	condition := builder.Equal(sectionIDColumn, sectionID)
	section := models.Section{
		ID: sectionID,
	}

	return update(r, &section, "UpdateSection", models.TableCoreSchPagSections, condition)
}

// DeleteSection deletes object from the database
func DeleteSection(r *http.Request) *Response {
	sectionID := chi.URLParam(r, "section_id")
	sectionIDColumn := fmt.Sprintf("%s.id", models.TableCoreSchPagSections)
	condition := builder.Equal(sectionIDColumn, sectionID)

	return remove(r, "DeleteSection", models.TableCoreSchPagSections, condition)
}
