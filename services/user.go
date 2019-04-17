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
	response := NewResponse()
	body, _ := ioutil.ReadAll(r.Body)
	user := &models.User{}
	err := json.Unmarshal(body, user)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorParsingRequest, "CreateUser unmarshal body", err.Error()))
		return response
	}

	id, err := db.InsertStruct(models.TableUsers, user)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorInsertingRecord, "CreateUser create", err.Error()))
		return response
	}
	user.ID = id

	response.Data = user

	return response
}

//LoadAllUsers return all instances from the object
func LoadAllUsers(r *http.Request) *Response {
	response := NewResponse()

	users := []models.User{}
	err := db.LoadStruct(models.TableUsers, &users, nil)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorLoadingData, "LoadAllUsers", err.Error()))
		return response
	}
	response.Data = users
	return response
}

//LoadUser return only one object from the database
func LoadUser(r *http.Request) *Response {
	response := NewResponse()
	userID := chi.URLParam(r, "user_id")
	user := &models.User{}
	err := db.LoadStruct(models.TableUsers, user, builder.Equal("users.id", userID))
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorLoadingData, "LoadUser", err.Error()))
		return response
	}
	response.Data = user
	return response
}

//UpdateUser updates object data in the database
func UpdateUser(r *http.Request) *Response {
	return nil
}

func DeleteUser(r *http.Request) *Response {
	return nil
}
