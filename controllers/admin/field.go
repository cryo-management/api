package admin

import (
	"net/http"

	"github.com/cryo-management/api/services"

	"github.com/go-chi/render"
)

func PostField(w http.ResponseWriter, r *http.Request) {
	response := services.CreateField(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

func GetAllFields(w http.ResponseWriter, r *http.Request) {
	response := services.LoadAllFields(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

func GetField(w http.ResponseWriter, r *http.Request) {
	response := services.LoadField(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

func UpdateField(w http.ResponseWriter, r *http.Request) {
	response := services.UpdateField(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

func DeleteField(w http.ResponseWriter, r *http.Request) {
	response := services.DeleteField(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}
