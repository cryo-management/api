package resources

import (
	"fmt"
	"net/http"
	"time"

	"github.com/andreluzz/go-sql-builder/builder"
	"github.com/andreluzz/go-sql-builder/db"
	"github.com/go-chi/chi"

	"github.com/cryo-management/api/models"
	"github.com/cryo-management/api/services"
)

// CreateGroup persists the request body creating a new object in the database
func CreateGroup(r *http.Request) *services.Response {
	group := models.Group{}

	return services.Create(r, &group, "CreateGroup", models.TableCoreGroups)
}

// LoadAllGroups return all instances from the object
func LoadAllGroups(r *http.Request) *services.Response {
	groups := []models.Group{}

	return services.Load(r, &groups, "LoadAllGroups", models.TableCoreGroups, nil)
}

// LoadGroup return only one object from the database
func LoadGroup(r *http.Request) *services.Response {
	group := models.Group{}
	groupID := chi.URLParam(r, "group_id")
	groupIDColumn := fmt.Sprintf("%s.id", models.TableCoreGroups)
	condition := builder.Equal(groupIDColumn, groupID)

	return services.Load(r, &group, "LoadGroup", models.TableCoreGroups, condition)
}

// UpdateGroup updates object data in the database
func UpdateGroup(r *http.Request) *services.Response {
	groupID := chi.URLParam(r, "group_id")
	groupIDColumn := fmt.Sprintf("%s.id", models.TableCoreGroups)
	condition := builder.Equal(groupIDColumn, groupID)
	group := models.Group{
		ID: groupID,
	}

	return services.Update(r, &group, "UpdateGroup", models.TableCoreGroups, condition)
}

// DeleteGroup deletes object from the database
func DeleteGroup(r *http.Request) *services.Response {
	groupID := chi.URLParam(r, "group_id")
	groupIDColumn := fmt.Sprintf("%s.id", models.TableCoreGroups)
	condition := builder.Equal(groupIDColumn, groupID)

	return services.Remove(r, "DeleteGroup", models.TableCoreGroups, condition)
}

// InsertUserInGroup persists the request creating a new object in the database
func InsertUserInGroup(r *http.Request) *services.Response {
	response := services.NewResponse()

	permissionGroupID := chi.URLParam(r, "group_id")
	permissionUserID := chi.URLParam(r, "user_id")

	userID := r.Header.Get("userID")
	now := time.Now()

	statemant := builder.Insert(
		models.TableCoreGroupsUsers,
		"group_id",
		"user_id",
		"created_by",
		"created_at",
		"updated_by",
		"updated_at",
	).Values(
		permissionGroupID,
		permissionUserID,
		userID,
		now,
		userID,
		now,
	)

	err := db.Exec(statemant)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, services.NewResponseError(services.ErrorInsertingRecord, "InsertUserInGroup", err.Error()))

		return response
	}

	return response
}

// LoadAllUsersByGroup return all instances from the object
func LoadAllUsersByGroup(r *http.Request) *services.Response {
	response := services.NewResponse()

	user := []models.User{}
	groupID := chi.URLParam(r, "group_id")

	statemant := builder.Select(
		"core_users.id",
		"core_users.first_name",
		"core_users.last_name",
		"core_users.email",
		"core_users.language_code",
		"core_users.active",
	).From(models.TableCoreUsers).Join(
		models.TableCoreGroupsUsers, "core_groups_users.user_id = core_users.id",
	).Where(
		builder.Equal("core_groups_users.group_id", groupID),
	)

	err := db.QueryStruct(statemant, &user)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, services.NewResponseError(services.ErrorLoadingData, "LoadAllUsersByGroup", err.Error()))

		return response
	}

	response.Data = user

	return response
}

// RemoveUserFromGroup deletes object from the database
func RemoveUserFromGroup(r *http.Request) *services.Response {
	response := services.NewResponse()

	groupID := chi.URLParam(r, "group_id")
	userID := chi.URLParam(r, "user_id")

	statemant := builder.Delete(models.TableCoreGroupsUsers).Where(
		builder.And(
			builder.Equal("group_id", groupID),
			builder.Equal("user_id", userID),
		),
	)

	err := db.Exec(statemant)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, services.NewResponseError(services.ErrorDeletingData, "RemoveUserFromGroup", err.Error()))

		return response
	}

	return response
}

// InsertPermission persists the request body creating a new object in the database
func InsertPermission(r *http.Request) *services.Response {
	permission := models.Permission{}

	return services.Create(r, &permission, "InsertPermission", models.TableCoreGrpPermissions)
}

// LoadAllPermissionsByGroup return all instances from the object
func LoadAllPermissionsByGroup(r *http.Request) *services.Response {
	permissions := []models.Permission{}
	groupID := chi.URLParam(r, "group_id")
	groupIDColumn := fmt.Sprintf("%s.group_id", models.TableCoreGrpPermissions)
	condition := builder.Equal(groupIDColumn, groupID)

	return services.Load(r, &permissions, "LoadAllPermissionsByGroup", models.TableCoreGrpPermissions, condition)
}

// RemovePermission deletes object from the database
func RemovePermission(r *http.Request) *services.Response {
	permissionID := chi.URLParam(r, "permission_id")
	permissionIDColumn := fmt.Sprintf("%s.id", models.TableCoreGrpPermissions)
	condition := builder.Equal(permissionIDColumn, permissionID)

	return services.Remove(r, "RemovePermission", models.TableCoreGrpPermissions, condition)
}
