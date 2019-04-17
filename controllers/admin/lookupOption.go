package admin

import (
	"net/http"

	"github.com/cryo-management/api/services"

	"github.com/go-chi/render"
)

func PostLookupOption(w http.ResponseWriter, r *http.Request) {
	response := services.CreateLookupOption(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

func GetAllLookupOptions(w http.ResponseWriter, r *http.Request) {
	response := services.LoadAllLookupOptions(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

func GetLookupOption(w http.ResponseWriter, r *http.Request) {
	response := services.LoadLookupOption(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

func UpdateLookupOption(w http.ResponseWriter, r *http.Request) {
	response := services.UpdateLookupOption(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

func DeleteLookupOption(w http.ResponseWriter, r *http.Request) {
	response := services.DeleteLookupOption(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}
