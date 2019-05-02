package resources

import (
	"fmt"
	"net/http"

	"github.com/andreluzz/go-sql-builder/builder"
	"github.com/go-chi/chi"

	"github.com/cryo-management/api/models"
	"github.com/cryo-management/api/services"
)

// CreateSection persists the request body creating a new object in the database
func CreateSection(r *http.Request) *services.Response {
	section := models.Section{}

	return services.Create(r, &section, "CreateSection", models.TableCoreSchPagSections)
}

// LoadAllSections return all instances from the object
func LoadAllSections(r *http.Request) *services.Response {
	sections := []models.Section{}
	pageID := chi.URLParam(r, "page_id")
	pageIDColumn := fmt.Sprintf("%s.page_id", models.TableCoreSchPagSections)
	condition := builder.Equal(pageIDColumn, pageID)

	return services.Load(r, &sections, "LoadAllSections", models.TableCoreSchPagSections, condition)
}

// LoadSection return only one object from the database
func LoadSection(r *http.Request) *services.Response {
	section := models.Section{}
	sectionID := chi.URLParam(r, "section_id")
	sectionIDColumn := fmt.Sprintf("%s.id", models.TableCoreSchPagSections)
	condition := builder.Equal(sectionIDColumn, sectionID)

	return services.Load(r, &section, "LoadSection", models.TableCoreSchPagSections, condition)
}

// UpdateSection updates object data in the database
func UpdateSection(r *http.Request) *services.Response {
	sectionID := chi.URLParam(r, "section_id")
	sectionIDColumn := fmt.Sprintf("%s.id", models.TableCoreSchPagSections)
	condition := builder.Equal(sectionIDColumn, sectionID)
	section := models.Section{
		ID: sectionID,
	}

	return services.Update(r, &section, "UpdateSection", models.TableCoreSchPagSections, condition)
}

// DeleteSection deletes object from the database
func DeleteSection(r *http.Request) *services.Response {
	sectionID := chi.URLParam(r, "section_id")
	sectionIDColumn := fmt.Sprintf("%s.id", models.TableCoreSchPagSections)
	condition := builder.Equal(sectionIDColumn, sectionID)

	return services.Remove(r, "DeleteSection", models.TableCoreSchPagSections, condition)
}
