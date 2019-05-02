package resources

import (
	"fmt"
	"net/http"

	"github.com/andreluzz/go-sql-builder/builder"
	"github.com/go-chi/chi"

	"github.com/cryo-management/api/models"
	"github.com/cryo-management/api/services"
)

// CreateLanguage persists the request body creating a new object in the database
func CreateLanguage(r *http.Request) *services.Response {
	language := models.Language{}

	return services.Create(r, &language, "CreateLanguage", models.TableCoreConfigLanguages)
}

// LoadAllLanguages return all instances from the object
func LoadAllLanguages(r *http.Request) *services.Response {
	languages := []models.Language{}

	return services.Load(r, &languages, "LoadAllLanguages", models.TableCoreConfigLanguages, nil)
}

// LoadLanguage return only one object from the database
func LoadLanguage(r *http.Request) *services.Response {
	language := models.Language{}
	languageID := chi.URLParam(r, "language_id")
	languageIDColumn := fmt.Sprintf("%s.id", models.TableCoreConfigLanguages)
	condition := builder.Equal(languageIDColumn, languageID)

	return services.Load(r, &language, "LoadLanguage", models.TableCoreConfigLanguages, condition)
}

// UpdateLanguage updates object data in the database
func UpdateLanguage(r *http.Request) *services.Response {
	languageID := chi.URLParam(r, "language_id")
	languageIDColumn := fmt.Sprintf("%s.id", models.TableCoreConfigLanguages)
	condition := builder.Equal(languageIDColumn, languageID)
	language := models.Language{
		ID: languageID,
	}

	return services.Update(r, &language, "UpdateLanguage", models.TableCoreConfigLanguages, condition)
}

// DeleteLanguage deletes object from the database
func DeleteLanguage(r *http.Request) *services.Response {
	languageID := chi.URLParam(r, "language_id")
	languageIDColumn := fmt.Sprintf("%s.id", models.TableCoreConfigLanguages)
	condition := builder.Equal(languageIDColumn, languageID)

	return services.Remove(r, "DeleteLanguage", models.TableCoreConfigLanguages, condition)
}
