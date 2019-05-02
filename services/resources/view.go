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

// CreateView persists the request body creating a new object in the database
func CreateView(r *http.Request) *services.Response {
	view := models.View{}

	return services.Create(r, &view, "CreateView", models.TableCoreSchViews)
}

// LoadAllViews return all instances from the object
func LoadAllViews(r *http.Request) *services.Response {
	views := []models.View{}
	schemaID := chi.URLParam(r, "schema_id")
	schemaIDColumn := fmt.Sprintf("%s.schema_id", models.TableCoreSchViews)
	condition := builder.Equal(schemaIDColumn, schemaID)

	return services.Load(r, &views, "LoadAllViews", models.TableCoreSchViews, condition)
}

// LoadView return only one object from the database
func LoadView(r *http.Request) *services.Response {
	view := models.View{}
	viewID := chi.URLParam(r, "view_id")
	viewIDColumn := fmt.Sprintf("%s.id", models.TableCoreSchViews)
	condition := builder.Equal(viewIDColumn, viewID)

	return services.Load(r, &view, "LoadView", models.TableCoreSchViews, condition)
}

// UpdateView updates object data in the database
func UpdateView(r *http.Request) *services.Response {
	viewID := chi.URLParam(r, "view_id")
	viewIDColumn := fmt.Sprintf("%s.id", models.TableCoreSchViews)
	condition := builder.Equal(viewIDColumn, viewID)
	view := models.View{
		ID: viewID,
	}

	return services.Update(r, &view, "UpdateView", models.TableCoreSchViews, condition)
}

// DeleteView deletes object from the database
func DeleteView(r *http.Request) *services.Response {
	viewID := chi.URLParam(r, "view_id")
	viewIDColumn := fmt.Sprintf("%s.id", models.TableCoreSchViews)
	condition := builder.Equal(viewIDColumn, viewID)

	return services.Remove(r, "DeleteView", models.TableCoreSchViews, condition)
}

// InsertPageInView persists the request creating a new object in the database
func InsertPageInView(r *http.Request) *services.Response {
	response := services.NewResponse()

	viewID := chi.URLParam(r, "view_id")
	pageID := chi.URLParam(r, "page_id")

	userID := r.Header.Get("userID")
	now := time.Now()

	statemant := builder.Insert(
		models.TableCoreViewsPages,
		"view_id",
		"page_id",
		"created_by",
		"created_at",
		"updated_by",
		"updated_at",
	).Values(
		viewID,
		pageID,
		userID,
		now,
		userID,
		now,
	)

	err := db.Exec(statemant)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, services.NewResponseError(services.ErrorInsertingRecord, "InsertPageInView", err.Error()))

		return response
	}

	return response
}

// LoadAllPagesByView return all instances from the object
func LoadAllPagesByView(r *http.Request) *services.Response {
	response := services.NewResponse()

	page := []models.Page{}
	viewID := chi.URLParam(r, "view_id")
	tblTranslationName := fmt.Sprintf("%s as %s_name", models.TableCoreTranslations, models.TableCoreTranslations)
	tblTranslationDescription := fmt.Sprintf("%s as %s_description", models.TableCoreTranslations, models.TableCoreTranslations)
	languageCode := r.Header.Get("Content-Language")

	statemant := builder.Select(
		"core_sch_pages.id",
		"core_sch_pages.code",
		"core_translations_name.value as name",
		"core_translations_description.value as description",
		"core_sch_pages.schema_id",
		"core_sch_pages.type",
		"core_sch_pages.active",
		"core_sch_pages.created_by",
		"core_sch_pages.created_at",
		"core_sch_pages.updated_by",
		"core_sch_pages.updated_at",
	).From(models.TableCoreSchPages).Join(
		tblTranslationName, "core_translations_name.structure_id = core_sch_pages.id and core_translations_name.structure_field = 'name'",
	).Join(
		tblTranslationDescription, "core_translations_description.structure_id = core_sch_pages.id and core_translations_description.structure_field = 'description'",
	).Join(
		models.TableCoreViewsPages, "core_views_pages.page_id = core_sch_pages.id",
	).Where(
		builder.And(
			builder.Equal("core_views_pages.view_id", viewID),
			builder.Equal("core_translations_name.language_code", languageCode),
			builder.Equal("core_translations_description.language_code", languageCode),
		),
	)

	err := db.QueryStruct(statemant, &page)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, services.NewResponseError(services.ErrorLoadingData, "LoadAllPagesByView", err.Error()))

		return response
	}

	response.Data = page

	return response
}

// RemovePageFromView deletes object from the database
func RemovePageFromView(r *http.Request) *services.Response {
	response := services.NewResponse()

	viewID := chi.URLParam(r, "view_id")
	pageID := chi.URLParam(r, "page_id")

	statemant := builder.Delete(models.TableCoreViewsPages).Where(
		builder.And(
			builder.Equal("view_id", viewID),
			builder.Equal("page_id", pageID),
		),
	)

	err := db.Exec(statemant)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, services.NewResponseError(services.ErrorDeletingData, "RemovePageFromView", err.Error()))

		return response
	}

	return response
}
