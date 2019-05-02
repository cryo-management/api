package resources

import (
	"fmt"
	"net/http"

	"github.com/andreluzz/go-sql-builder/builder"
	"github.com/andreluzz/go-sql-builder/db"
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

// LoadAllGroupsByUser return all instances from the object
func LoadAllGroupsByUser(r *http.Request) *services.Response {
	response := services.NewResponse()

	group := []models.Group{}
	userID := chi.URLParam(r, "user_id")
	tblTranslationName := fmt.Sprintf("%s as %s_name", models.TableCoreTranslations, models.TableCoreTranslations)
	tblTranslationDescription := fmt.Sprintf("%s as %s_description", models.TableCoreTranslations, models.TableCoreTranslations)
	languageCode := r.Header.Get("Content-Language")

	statemant := builder.Select(
		"core_groups.id",
		"core_translations_name.value as name",
		"core_translations_description.value as description",
		"core_groups.code",
	).From(models.TableCoreGroups).Join(
		tblTranslationName, "core_translations_name.structure_id = core_groups.id and core_translations_name.structure_field = 'name'",
	).Join(
		tblTranslationDescription, "core_translations_description.structure_id = core_groups.id and core_translations_description.structure_field = 'description'",
	).Join(
		models.TableCoreGroupsUsers, "core_groups_users.group_id = core_groups.id",
	).Where(
		builder.And(
			builder.Equal("core_groups_users.user_id", userID),
			builder.Equal("core_translations_name.language_code", languageCode),
			builder.Equal("core_translations_description.language_code", languageCode),
		),
	)

	err := db.QueryStruct(statemant, &group)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, services.NewResponseError(services.ErrorLoadingData, "LoadAllGroupsByUser", err.Error()))

		return response
	}

	response.Data = group

	return response
}
