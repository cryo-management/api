package admin

import (
	"net/http"

	"github.com/cryo-management/api/services"

	"github.com/go-chi/render"
)

func PostLookup(w http.ResponseWriter, r *http.Request) {
	response := services.CreateLookup(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

func GetAllLookups(w http.ResponseWriter, r *http.Request) {
	response := services.LoadAllLookups(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

func GetLookup(w http.ResponseWriter, r *http.Request) {
	response := services.LoadLookup(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

func UpdateLookup(w http.ResponseWriter, r *http.Request) {
	response := services.UpdateLookup(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

func DeleteLookup(w http.ResponseWriter, r *http.Request) {
	response := services.DeleteLookup(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

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
