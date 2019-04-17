package admin

import (
	"net/http"

	"github.com/cryo-management/api/services"

	"github.com/go-chi/render"
)

func PostGroup(w http.ResponseWriter, r *http.Request) {
	response := services.CreateGroup(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

func GetAllGroups(w http.ResponseWriter, r *http.Request) {
	response := services.LoadAllGroups(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

func GetGroup(w http.ResponseWriter, r *http.Request) {
	response := services.LoadGroup(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

func UpdateGroup(w http.ResponseWriter, r *http.Request) {
	response := services.UpdateGroup(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

func DeleteGroup(w http.ResponseWriter, r *http.Request) {
	response := services.DeleteGroup(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

func PostGroupUser(w http.ResponseWriter, r *http.Request) {
	render.Status(r, 500)
	render.JSON(w, r, "")
}

func DeleteGroupUser(w http.ResponseWriter, r *http.Request) {
	render.Status(r, 500)
	render.JSON(w, r, "")
}

func PostGroupPermission(w http.ResponseWriter, r *http.Request) {
	render.Status(r, 500)
	render.JSON(w, r, "")
}

func DeleteGroupPermission(w http.ResponseWriter, r *http.Request) {
	render.Status(r, 500)
	render.JSON(w, r, "")
}
