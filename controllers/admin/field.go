package admin

import (
	"net/http"

	"github.com/cryo-management/api/services"

	"github.com/go-chi/render"
)

// PostField sends the request to service creating a new field
func PostField(w http.ResponseWriter, r *http.Request) {
	response := services.CreateField(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// GetAllFields return all field instances from the service
func GetAllFields(w http.ResponseWriter, r *http.Request) {
	response := services.LoadAllFields(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// GetField return only one field from the service
func GetField(w http.ResponseWriter, r *http.Request) {
	response := services.LoadField(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// UpdateField sends the request to service updating a field
func UpdateField(w http.ResponseWriter, r *http.Request) {
	response := services.UpdateField(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// DeleteField sends the request to service deleting a field
func DeleteField(w http.ResponseWriter, r *http.Request) {
	response := services.DeleteField(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}
