package services

import (
	"net/http"

	"github.com/andreluzz/go-sql-builder/builder"
	"github.com/andreluzz/go-sql-builder/db"
	"github.com/go-chi/chi"

	"github.com/cryo-management/api/models"
)

// CreateGroup persists the request body creating a new object in the database
func CreateGroup(r *http.Request) *Response {
	group := models.Group{}

	return create(r, &group, "CreateGroup", models.TableGroups)
}

// LoadAllGroups return all instances from the object
func LoadAllGroups(r *http.Request) *Response {
	groups := []models.Group{}

	return load(r, &groups, "LoadAllGroups", models.TableGroups, nil)
}

// LoadGroup return only one object from the database
func LoadGroup(r *http.Request) *Response {
	group := models.Group{}
	groupID := chi.URLParam(r, "group_id")
	condition := builder.Equal("groups.id", groupID)

	return load(r, &group, "LoadGroup", models.TableGroups, condition)
}

// UpdateGroup updates object data in the database
func UpdateGroup(r *http.Request) *Response {
	groupID := chi.URLParam(r, "group_id")
	condition := builder.Equal("groups.id", groupID)
	group := models.Group{
		ID: groupID,
	}

	return update(r, &group, "UpdateGroup", models.TableGroups, condition)
}

// DeleteGroup deletes object from the database
func DeleteGroup(r *http.Request) *Response {
	groupID := chi.URLParam(r, "group_id")
	condition := builder.Equal("groups.id", groupID)

	return remove(r, "DeleteGroup", models.TableGroups, condition)
}

// InsertUserInGroup persists the request creating a new object in the database
func InsertUserInGroup(r *http.Request) *Response {
	response := NewResponse()

	groupID := chi.URLParam(r, "group_id")
	userID := chi.URLParam(r, "user_id")

	statemant := builder.Insert(models.TableGroupsUsers, "group_id", "user_id").Values(groupID, userID)

	err := db.Exec(statemant)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorInsertingRecord, "InsertUserInGroup", err.Error()))

		return response
	}

	return response
}

// LoadAllUsersByGroup return all instances from the object
func LoadAllUsersByGroup(r *http.Request) *Response {
	response := NewResponse()

	user := []models.User{}
	groupID := chi.URLParam(r, "group_id")

	statemant := builder.Select(
		"users.id", "users.first_name", "users.last_name", "users.email", "users.language", "users.active",
	).From(models.TableUsers).Join(models.TableGroupsUsers, "groups_users.user_id = users.id").Where(
		builder.Equal("groups_users.group_id", groupID),
	)

	err := db.QueryStruct(statemant, &user)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorLoadingData, "LoadAllUsersByGroup", err.Error()))

		return response
	}

	response.Data = user

	return response
}

// RemoveUserFromGroup deletes object from the database
func RemoveUserFromGroup(r *http.Request) *Response {
	response := NewResponse()

	groupID := chi.URLParam(r, "group_id")
	userID := chi.URLParam(r, "user_id")

	statemant := builder.Delete(models.TableGroupsUsers).Where(
		builder.And(
			builder.Equal("group_id", groupID),
			builder.Equal("user_id", userID),
		),
	)

	err := db.Exec(statemant)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, NewResponseError(ErrorDeletingData, "RemoveUserFromGroup", err.Error()))

		return response
	}

	return response
}

// InsertPermission persists the request body creating a new object in the database
func InsertPermission(r *http.Request) *Response {
	permission := models.Permission{}

	return create(r, &permission, "InsertPermission", models.TableGroupsPermissions)
}

// LoadAllPermissionsByGroup return all instances from the object
func LoadAllPermissionsByGroup(r *http.Request) *Response {
	permissions := []models.Permission{}
	groupID := chi.URLParam(r, "group_id")
	condition := builder.Equal("groups_permissions.group_id", groupID)

	return load(r, &permissions, "LoadAllPermissionsByGroup", models.TableGroupsPermissions, condition)
}

// RemovePermission deletes object from the database
func RemovePermission(r *http.Request) *Response {
	permissionID := chi.URLParam(r, "permission_id")
	condition := builder.Equal("groups_permissions.id", permissionID)

	return remove(r, "RemovePermission", models.TableGroupsPermissions, condition)
}
