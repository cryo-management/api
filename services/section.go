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

// CreateSectionStructure persists the request body creating a new object in the database
func CreateSectionStructure(r *http.Request) *Response {
	sectionStructure := models.SectionStructure{}

	return create(r, &sectionStructure, "CreateSectionStructure", models.TableCoreSchPagSecStructures)
}

// LoadAllSectionStructures return all instances from the object
func LoadAllSectionStructures(r *http.Request) *Response {
	sectionStructures := []models.SectionStructure{}
	sectionID := chi.URLParam(r, "section_id")
	sectionIDColumn := fmt.Sprintf("%s.section_id", models.TableCoreSchPagSecStructures)
	condition := builder.Equal(sectionIDColumn, sectionID)

	return load(r, &sectionStructures, "LoadAllSectionStructures", models.TableCoreSchPagSecStructures, condition)
}

// LoadSectionStructure return only one object from the database
func LoadSectionStructure(r *http.Request) *Response {
	sectionStructure := models.SectionStructure{}
	sectionStructureID := chi.URLParam(r, "section_structure_id")
	sectionStructureIDColumn := fmt.Sprintf("%s.id", models.TableCoreSchPagSecStructures)
	condition := builder.Equal(sectionStructureIDColumn, sectionStructureID)

	return load(r, &sectionStructure, "LoadSectionStructure", models.TableCoreSchPagSecStructures, condition)
}

// UpdateSectionStructure updates object data in the database
func UpdateSectionStructure(r *http.Request) *Response {
	sectionStructureID := chi.URLParam(r, "section_structure_id")
	sectionStructureIDColumn := fmt.Sprintf("%s.id", models.TableCoreSchPagSecStructures)
	condition := builder.Equal(sectionStructureIDColumn, sectionStructureID)
	sectionStructure := models.SectionStructure{
		ID: sectionStructureID,
	}

	return update(r, &sectionStructure, "UpdateSectionStructure", models.TableCoreSchPagSecStructures, condition)
}

// DeleteSectionStructure deletes object from the database
func DeleteSectionStructure(r *http.Request) *Response {
	sectionStructureID := chi.URLParam(r, "section_structure_id")
	sectionStructureIDColumn := fmt.Sprintf("%s.id", models.TableCoreSchPagSecStructures)
	condition := builder.Equal(sectionStructureIDColumn, sectionStructureID)

	return remove(r, "DeleteSectionStructure", models.TableCoreSchPagSecStructures, condition)
}
