package services

import (
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
	viewID := chi.URLParam(r, "view_id")
	condition := builder.Equal("sections.view_id", viewID)

	return load(r, &sections, "LoadAllSections", models.TableCoreSchPagSections, condition)
}

// LoadSection return only one object from the database
func LoadSection(r *http.Request) *Response {
	section := models.Section{}
	sectionID := chi.URLParam(r, "section_id")
	condition := builder.Equal("sections.id", sectionID)

	return load(r, &section, "LoadSection", models.TableCoreSchPagSections, condition)
}

// UpdateSection updates object data in the database
func UpdateSection(r *http.Request) *Response {
	sectionID := chi.URLParam(r, "section_id")
	condition := builder.Equal("sections.id", sectionID)
	section := models.Section{
		ID: sectionID,
	}

	return update(r, &section, "UpdateSection", models.TableCoreSchPagSections, condition)
}

// DeleteSection deletes object from the database
func DeleteSection(r *http.Request) *Response {
	sectionID := chi.URLParam(r, "section_id")
	condition := builder.Equal("sections.id", sectionID)

	return remove(r, "DeleteSection", models.TableCoreSchPagSections, condition)
}

// CreateSectionStructure persists the request body creating a new object in the database
func CreateSectionStructure(r *http.Request) *Response {
	sectionStructure := models.SectionStructure{}

	return create(r, &sectionStructure, "CreateSectionStructure", models.TableCoreSchPagSecStructures)
}

// LoadAllSectionStructures return all instances from the object
func LoadAllSectionStructures(r *http.Request) *Response {
	sectionStructures := []models.SectionStructure{}
	viewID := chi.URLParam(r, "view_id")
	condition := builder.Equal("sectionStructures.view_id", viewID)

	return load(r, &sectionStructures, "LoadAllSectionStructures", models.TableCoreSchPagSecStructures, condition)
}

// LoadSectionStructure return only one object from the database
func LoadSectionStructure(r *http.Request) *Response {
	sectionStructure := models.SectionStructure{}
	sectionStructureID := chi.URLParam(r, "sectionStructure_id")
	condition := builder.Equal("sectionStructures.id", sectionStructureID)

	return load(r, &sectionStructure, "LoadSectionStructure", models.TableCoreSchPagSecStructures, condition)
}

// UpdateSectionStructure updates object data in the database
func UpdateSectionStructure(r *http.Request) *Response {
	sectionStructureID := chi.URLParam(r, "sectionStructure_id")
	condition := builder.Equal("sectionStructures.id", sectionStructureID)
	sectionStructure := models.SectionStructure{
		ID: sectionStructureID,
	}

	return update(r, &sectionStructure, "UpdateSectionStructure", models.TableCoreSchPagSecStructures, condition)
}

// DeleteSectionStructure deletes object from the database
func DeleteSectionStructure(r *http.Request) *Response {
	sectionStructureID := chi.URLParam(r, "sectionStructure_id")
	condition := builder.Equal("sectionStructures.id", sectionStructureID)

	return remove(r, "DeleteSectionStructure", models.TableCoreSchPagSecStructures, condition)
}
