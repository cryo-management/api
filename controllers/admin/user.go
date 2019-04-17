package admin

import (
	"net/http"

	"github.com/cryo-management/api/services"

	"github.com/go-chi/render"
)

func PostUser(w http.ResponseWriter, r *http.Request) {
	response := services.CreateUser(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	response := services.LoadAllUsers(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	response := services.LoadUser(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	response := services.UpdateUser(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	response := services.DeleteUser(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}
