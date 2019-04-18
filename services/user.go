package services

import (
	"encoding/json"
	"fmt"
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

	return load(r, &user, "LoadUser", models.TableUsers, condition)
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

	return remove(r, "DeleteUser", models.TableUsers, condition)
}

//LoadAllGroupsByUser return all instances from the object
func LoadAllGroupsByUser(r *http.Request) *Response {
	response := NewResponse()

	group := []models.Group{}
	userID := chi.URLParam(r, "user_id")
	tblTranslationName := fmt.Sprintf("%s as %s_name", models.TableTranslations, models.TableTranslations)
	tblTranslationDescription := fmt.Sprintf("%s as %s_description", models.TableTranslations, models.TableTranslations)
	languageCode := r.Header.Get("languageCode")

	statemant := builder.Select(
		"groups.id", "translations_name.value as name", "translations_description.value as description", "groups.code",
	).From(models.TableGroups).Join(
		tblTranslationName, "translations_name.structure_id = groups.id and translations_name.structure_field = 'name'",
	).Join(
		tblTranslationDescription, "translations_description.structure_id = groups.id and translations_description.structure_field = 'description'",
	).Join(
		models.TableGroupsUsers, "groups_users.group_id = groups.id",
	).Where(
		builder.And(
			builder.Equal("groups_users.user_id", userID),
			builder.Equal("translations_name.language_code", languageCode),
			builder.Equal("translations_description.language_code", languageCode),
		),
	)

	err := db.QueryStruct(statemant, &group)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorLoadingData, "LoadAllGroupsByUser", err.Error()))

		return response
	}

	response.Data = group

	return response
}
