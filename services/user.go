package services

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/andreluzz/go-sql-builder/builder"
	"github.com/andreluzz/go-sql-builder/db"
	"github.com/go-chi/chi"

	"github.com/cryo-management/api/models"
)

//CreateUser persists the request body creating a new object in the database
func CreateUser(r *http.Request) *Response {
	user := models.User{}

	return create(r, &user, "CreateUser", models.TableUsers)
}

//LoadAllUsers return all instances from the object
func LoadAllUsers(r *http.Request) *Response {
	users := []models.User{}

	return load(r, &users, "LoadAllUsers", models.TableUsers, nil)
}

//LoadUser return only one object from the database
func LoadUser(r *http.Request) *Response {
	user := models.User{}
	userID := chi.URLParam(r, "user_id")
	condition := builder.Equal("users.id", userID)

	return load(r, &user, "LoadAUser", models.TableUsers, condition)
}

//UpdateUser updates object data in the database
func UpdateUser(r *http.Request) *Response {
	response := NewResponse()
	userID := chi.URLParam(r, "user_id")
	user := &models.User{}
	body, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(body, user)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorParsingRequest, "UpdateUser unmarshal body", err.Error()))

		return response
	}

	condition := builder.Equal("users.id", userID)
	columns := getColumnsFromBody(body)

	err = db.UpdateStruct(models.TableUsers, user, condition, columns...)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorInsertingRecord, "UpdateUser", err.Error()))

		return response
	}

	return response
}

//DeleteUser deletes object from the database
func DeleteUser(r *http.Request) *Response {
	userID := chi.URLParam(r, "user_id")
	condition := builder.Equal("users.id", userID)

	return delete(r, "DeleteUser", models.TableUsers, condition)
}
