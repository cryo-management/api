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

//CreateUser persists the request body creating a new user in the database
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

func LoadAllUsers(r *http.Request) *Response {
	response := NewResponse()

	users := []models.User{}
	jsonBytes, err := db.LoadStruct(models.TableUsers, users, nil)
	json.Unmarshal(jsonBytes, &users)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorLoadingData, "LoadAllUsers", err.Error()))
		return response
	}
	response.Data = users
	return response
}

func LoadUser(r *http.Request) *Response {
	response := NewResponse()
	userID := chi.URLParam(r, "user_id")
	user := &models.User{}
	jsonBytes, err := db.LoadStruct(models.TableUsers, user, builder.Equal("users.id", userID))
	json.Unmarshal(jsonBytes, user)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorLoadingData, "LoadUser", err.Error()))
		return response
	}
	response.Data = user
	return response
}

func UpdateUser(r *http.Request) *Response {
	return nil
}

func DeleteUser(r *http.Request) *Response {
	return nil
}
