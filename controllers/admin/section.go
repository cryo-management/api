package admin

import (
	"net/http"

	"github.com/cryo-management/api/services"

	"github.com/go-chi/render"
)

// PostSection sends the request to service creating a new section
func PostSection(w http.ResponseWriter, r *http.Request) {
	response := services.CreateSection(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// GetAllSections return all section instances from the service
func GetAllSections(w http.ResponseWriter, r *http.Request) {
	response := services.LoadAllSections(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// GetSection return only one section from the service
func GetSection(w http.ResponseWriter, r *http.Request) {
	response := services.LoadSection(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// UpdateSection sends the request to service updating a section
func UpdateSection(w http.ResponseWriter, r *http.Request) {
	response := services.UpdateSection(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// DeleteSection sends the request to service deleting a section
func DeleteSection(w http.ResponseWriter, r *http.Request) {
	response := services.DeleteSection(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// PostSectionStructure sends the request to service creating a new sectionStructure
func PostSectionStructure(w http.ResponseWriter, r *http.Request) {
	response := services.CreateSectionStructure(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// GetAllSectionStructures return all sectionStructure instances from the service
func GetAllSectionStructures(w http.ResponseWriter, r *http.Request) {
	response := services.LoadAllSectionStructures(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// GetSectionStructure return only one sectionStructure from the service
func GetSectionStructure(w http.ResponseWriter, r *http.Request) {
	response := services.LoadSectionStructure(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// UpdateSectionStructure sends the request to service updating a sectionStructure
func UpdateSectionStructure(w http.ResponseWriter, r *http.Request) {
	response := services.UpdateSectionStructure(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// DeleteSectionStructure sends the request to service deleting a sectionStructure
func DeleteSectionStructure(w http.ResponseWriter, r *http.Request) {
	response := services.DeleteSectionStructure(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}
