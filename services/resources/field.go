package resources

import (
	"fmt"
	"net/http"

	"github.com/andreluzz/go-sql-builder/builder"
	"github.com/go-chi/chi"

	"github.com/cryo-management/api/models"
	"github.com/cryo-management/api/services"
)

// CreateField persists the request body creating a new object in the database
func CreateField(r *http.Request) *services.Response {
	field := models.Field{}

	return services.Create(r, &field, "CreateField", models.TableCoreSchFields)
}

// LoadAllFields return all instances from the object
func LoadAllFields(r *http.Request) *services.Response {
	fields := []models.Field{}
	schemaID := chi.URLParam(r, "schema_id")
	schemaIDColumn := fmt.Sprintf("%s.schema_id", models.TableCoreSchFields)
	condition := builder.Equal(schemaIDColumn, schemaID)

	return services.Load(r, &fields, "LoadAllFields", models.TableCoreSchFields, condition)
}

// LoadField return only one object from the database
func LoadField(r *http.Request) *services.Response {
	field := models.Field{}
	fieldID := chi.URLParam(r, "field_id")
	fieldIDColumn := fmt.Sprintf("%s.id", models.TableCoreSchFields)
	condition := builder.Equal(fieldIDColumn, fieldID)

	return services.Load(r, &field, "LoadField", models.TableCoreSchFields, condition)
}

// UpdateField updates object data in the database
func UpdateField(r *http.Request) *services.Response {
	fieldID := chi.URLParam(r, "field_id")
	fieldIDColumn := fmt.Sprintf("%s.id", models.TableCoreSchFields)
	condition := builder.Equal(fieldIDColumn, fieldID)
	field := models.Field{
		ID: fieldID,
	}

	return services.Update(r, &field, "UpdateField", models.TableCoreSchFields, condition)
}

// DeleteField deletes object from the database
func DeleteField(r *http.Request) *services.Response {
	fieldID := chi.URLParam(r, "field_id")
	fieldIDColumn := fmt.Sprintf("%s.id", models.TableCoreSchFields)
	condition := builder.Equal(fieldIDColumn, fieldID)

	return services.Remove(r, "DeleteField", models.TableCoreSchFields, condition)
}

// CreateFieldValidation persists the request body creating a new object in the database
func CreateFieldValidation(r *http.Request) *services.Response {
	fieldValidation := models.FieldValidation{}

	return services.Create(r, &fieldValidation, "CreateFieldValidation", models.TableCoreSchFldValidations)
}

// LoadAllFieldValidations return all instances from the object
func LoadAllFieldValidations(r *http.Request) *services.Response {
	fieldValidations := []models.FieldValidation{}
	fieldID := chi.URLParam(r, "field_id")
	fieldIDColumn := fmt.Sprintf("%s.field_id", models.TableCoreSchFldValidations)
	condition := builder.Equal(fieldIDColumn, fieldID)

	return services.Load(r, &fieldValidations, "LoadAllFieldValidations", models.TableCoreSchFldValidations, condition)
}

// LoadFieldValidation return only one object from the database
func LoadFieldValidation(r *http.Request) *services.Response {
	fieldValidation := models.FieldValidation{}
	fieldValidationID := chi.URLParam(r, "field_validation_id")
	fieldValidationIDColumn := fmt.Sprintf("%s.id", models.TableCoreSchFldValidations)
	condition := builder.Equal(fieldValidationIDColumn, fieldValidationID)

	return services.Load(r, &fieldValidation, "LoadFieldValidation", models.TableCoreSchFldValidations, condition)
}

// UpdateFieldValidation updates object data in the database
func UpdateFieldValidation(r *http.Request) *services.Response {
	fieldValidationID := chi.URLParam(r, "field_validation_id")
	fieldValidationIDColumn := fmt.Sprintf("%s.id", models.TableCoreSchFldValidations)
	condition := builder.Equal(fieldValidationIDColumn, fieldValidationID)
	fieldValidation := models.FieldValidation{
		ID: fieldValidationID,
	}

	return services.Update(r, &fieldValidation, "UpdateFieldValidation", models.TableCoreSchFldValidations, condition)
}

// DeleteFieldValidation deletes object from the database
func DeleteFieldValidation(r *http.Request) *services.Response {
	fieldValidationID := chi.URLParam(r, "field_validation_id")
	fieldValidationIDColumn := fmt.Sprintf("%s.id", models.TableCoreSchFldValidations)
	condition := builder.Equal(fieldValidationIDColumn, fieldValidationID)

	return services.Remove(r, "DeleteFieldValidation", models.TableCoreSchFldValidations, condition)
}
