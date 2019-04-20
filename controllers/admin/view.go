package admin

import (
	"net/http"

	"github.com/cryo-management/api/services"

	"github.com/go-chi/render"
)

// PostView sends the request to service creating a new view
func PostView(w http.ResponseWriter, r *http.Request) {
	response := services.CreateView(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// GetAllViews return all view instances from the service
func GetAllViews(w http.ResponseWriter, r *http.Request) {
	response := services.LoadAllViews(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// GetView return only one view from the service
func GetView(w http.ResponseWriter, r *http.Request) {
	response := services.LoadView(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// UpdateView sends the request to service updating a view
func UpdateView(w http.ResponseWriter, r *http.Request) {
	response := services.UpdateView(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// DeleteView sends the request to service deleting a view
func DeleteView(w http.ResponseWriter, r *http.Request) {
	response := services.DeleteView(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}
