package resources

import (
	"fmt"
	"net/http"

	"github.com/andreluzz/go-sql-builder/builder"
	"github.com/go-chi/chi"

	"github.com/cryo-management/api/models"
	"github.com/cryo-management/api/services"
)

// CreateUser persists the request body creating a new object in the database
func CreateUser(r *http.Request) *services.Response {
	user := models.User{}

	return services.Create(r, &user, "CreateUser", models.TableCoreUsers)
}

// LoadAllUsers return all instances from the object
func LoadAllUsers(r *http.Request) *services.Response {
	users := []models.User{}

	return services.Load(r, &users, "LoadAllUsers", models.TableCoreUsers, nil)
}

// LoadUser return only one object from the database
func LoadUser(r *http.Request) *services.Response {
	user := models.User{}
	userID := chi.URLParam(r, "user_id")
	userIDColumn := fmt.Sprintf("%s.id", models.TableCoreUsers)
	condition := builder.Equal(userIDColumn, userID)

	return services.Load(r, &user, "LoadUser", models.TableCoreUsers, condition)
}

// UpdateUser updates object data in the database
func UpdateUser(r *http.Request) *services.Response {
	userID := chi.URLParam(r, "user_id")
	userIDColumn := fmt.Sprintf("%s.id", models.TableCoreUsers)
	condition := builder.Equal(userIDColumn, userID)
	user := models.User{
		ID: userID,
	}

	return services.Update(r, &user, "UpdateUser", models.TableCoreUsers, condition)
}

// DeleteUser deletes object from the database
func DeleteUser(r *http.Request) *services.Response {
	userID := chi.URLParam(r, "user_id")
	userIDColumn := fmt.Sprintf("%s.id", models.TableCoreUsers)
	condition := builder.Equal(userIDColumn, userID)

	return services.Remove(r, "DeleteUser", models.TableCoreUsers, condition)
}

// LoadAllUsersByGroup return all instances from the object
func LoadAllUsersByGroup(r *http.Request) *services.Response {
	viewGroupUsers := []models.ViewGroupUser{}
	groupID := chi.URLParam(r, "group_id")
	groupIDColumn := fmt.Sprintf("%s.group_id", models.ViewCoreGroupUsers)
	condition := builder.Equal(groupIDColumn, groupID)

	return services.Load(r, &viewGroupUsers, "LoadAllUsersByGroup", models.ViewCoreGroupUsers, condition)
}
