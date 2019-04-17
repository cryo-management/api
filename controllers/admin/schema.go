package admin

import (
	"net/http"

	"github.com/cryo-management/api/services"

	"github.com/go-chi/render"
)

func PostSchema(w http.ResponseWriter, r *http.Request) {
	response := services.CreateSchema(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

func GetAllSchemas(w http.ResponseWriter, r *http.Request) {
	response := services.LoadAllSchemas(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

func GetSchema(w http.ResponseWriter, r *http.Request) {
	response := services.LoadSchema(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

func UpdateSchema(w http.ResponseWriter, r *http.Request) {
	response := services.UpdateSchema(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

func DeleteSchema(w http.ResponseWriter, r *http.Request) {
	response := services.DeleteSchema(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}
