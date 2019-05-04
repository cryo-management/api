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

// CreateSchema persists the request body creating a new object in the database
func CreateSchema(r *http.Request) *services.Response {
	schema := models.Schema{}

	return services.Create(r, &schema, "CreateSchema", models.TableCoreSchemas)
}

// LoadAllSchemas return all instances from the object
func LoadAllSchemas(r *http.Request) *services.Response {
	schemas := []models.Schema{}

	return services.Load(r, &schemas, "LoadAllSchemas", models.TableCoreSchemas, nil)
}

// LoadSchema return only one object from the database
func LoadSchema(r *http.Request) *services.Response {
	schema := models.Schema{}
	schemaID := chi.URLParam(r, "schema_id")
	schemaIDColumn := fmt.Sprintf("%s.id", models.TableCoreSchemas)
	condition := builder.Equal(schemaIDColumn, schemaID)

	return services.Load(r, &schema, "LoadSchema", models.TableCoreSchemas, condition)
}

// UpdateSchema updates object data in the database
func UpdateSchema(r *http.Request) *services.Response {
	schemaID := chi.URLParam(r, "schema_id")
	schemaIDColumn := fmt.Sprintf("%s.id", models.TableCoreSchemas)
	condition := builder.Equal(schemaIDColumn, schemaID)
	schema := models.Schema{
		ID: schemaID,
	}

	return services.Update(r, &schema, "UpdateSchema", models.TableCoreSchemas, condition)
}

// DeleteSchema deletes object from the database
func DeleteSchema(r *http.Request) *services.Response {
	schemaID := chi.URLParam(r, "schema_id")
	schemaIDColumn := fmt.Sprintf("%s.id", models.TableCoreSchemas)
	condition := builder.Equal(schemaIDColumn, schemaID)

	return services.Remove(r, "DeleteSchema", models.TableCoreSchemas, condition)
}

// InsertModuleInSchema persists the request creating a new object in the database
func InsertModuleInSchema(r *http.Request) *services.Response {
	response := services.NewResponse()

	schemaID := chi.URLParam(r, "schema_id")
	moduleID := chi.URLParam(r, "module_id")

	userID := r.Header.Get("userID")
	now := time.Now()

	statemant := builder.Insert(
		models.TableCoreSchemasModels,
		"schema_id",
		"module_id",
		"created_by",
		"created_at",
		"updated_by",
		"updated_at",
	).Values(
		schemaID,
		moduleID,
		userID,
		now,
		userID,
		now,
	)

	err := db.Exec(statemant)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, services.NewResponseError(services.ErrorInsertingRecord, "InsertModuleSchema", err.Error()))

		return response
	}

	return response
}

// LoadAllModulesBySchema return all instances from the object
func LoadAllModulesBySchema(r *http.Request) *services.Response {
	response := services.NewResponse()

	module := []models.Schema{}
	schemaID := chi.URLParam(r, "schema_id")
	tblTranslationName := fmt.Sprintf("%s as %s_name", models.TableCoreTranslations, models.TableCoreTranslations)
	tblTranslationDescription := fmt.Sprintf("%s as %s_description", models.TableCoreTranslations, models.TableCoreTranslations)
	languageCode := r.Header.Get("Content-Language")

	statemant := builder.Select(
		"core_schemas.id",
		"core_schemas.code",
		"core_translations_name.value as name",
		"core_translations_description.value as description",
		"core_schemas.module",
		"core_schemas.active",
		"core_schemas.created_by",
		"core_schemas.created_at",
		"core_schemas.updated_by",
		"core_schemas.updated_at",
	).From(models.TableCoreSchemas).Join(
		tblTranslationName, "core_translations_name.structure_id = core_schemas.id and core_translations_name.structure_field = 'name'",
	).Join(
		tblTranslationDescription, "core_translations_description.structure_id = core_schemas.id and core_translations_description.structure_field = 'description'",
	).Join(
		models.TableCoreSchemasModels, "core_schemas_modules.module_id = core_schemas.id",
	).Where(
		builder.And(
			builder.Equal("core_schemas_modules.schema_id", schemaID),
			builder.Equal("core_translations_name.language_code", languageCode),
			builder.Equal("core_translations_description.language_code", languageCode),
		),
	)

	err := db.QueryStruct(statemant, &module)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, services.NewResponseError(services.ErrorLoadingData, "LoadAllModulesBySchema", err.Error()))

		return response
	}

	response.Data = module

	return response
}

// RemoveModuleFromSchema deletes object from the database
func RemoveModuleFromSchema(r *http.Request) *services.Response {
	response := services.NewResponse()

	schemaID := chi.URLParam(r, "schema_id")
	moduleID := chi.URLParam(r, "module_id")

	statemant := builder.Delete(models.TableCoreSchemasModels).Where(
		builder.And(
			builder.Equal("schema_id", schemaID),
			builder.Equal("module_id", moduleID),
		),
	)

	err := db.Exec(statemant)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, services.NewResponseError(services.ErrorDeletingData, "RemoveModuleFromSchema", err.Error()))

		return response
	}

	return response
}
