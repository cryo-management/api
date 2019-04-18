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
	response := services.InsertUserInGroup(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

func GetAllUsersByGroup(w http.ResponseWriter, r *http.Request) {
	response := services.LoadAllUsersByGroup(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

func DeleteGroupUser(w http.ResponseWriter, r *http.Request) {
	response := services.RemoveUserFromGroup(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

func PostGroupPermission(w http.ResponseWriter, r *http.Request) {
	response := services.InsertPermission(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

func GetAllPermissionsByGroup(w http.ResponseWriter, r *http.Request) {
	response := services.LoadAllPermissionsByGroup(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

func DeleteGroupPermission(w http.ResponseWriter, r *http.Request) {
	response := services.RemovePermission(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}
