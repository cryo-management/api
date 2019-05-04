package models

import (
	"fmt"
	"reflect"
	"time"

	"github.com/andreluzz/go-sql-builder/builder"
	"github.com/andreluzz/go-sql-builder/db"
)

// Translation defines the struct of this object
type Translation struct {
	ID             string    `json:"id" sql:"id" pk:"true"`
	StructureID    string    `json:"structure_id" sql:"structure_id" fk:"true"`
	StructureType  string    `json:"structure_type" sql:"structure_type"`
	StructureField string    `json:"structure_field" sql:"structure_field"`
	Value          string    `json:"value" sql:"value"`
	LanguageCode   string    `json:"language_code" sql:"language_code"`
	Replicated     bool      `json:"replicated" sql:"replicated"`
	CreatedBy      string    `json:"created_by" sql:"created_by"`
	CreatedByUser  *User     `json:"created_by_user" table:"core_users" alias:"created_by_user" on:"created_by_user.id = core_translations.created_by"`
	CreatedAt      time.Time `json:"created_at" sql:"created_at"`
	UpdatedBy      string    `json:"updated_by" sql:"updated_by"`
	UpdatedByUser  *User     `json:"updated_by_user" table:"core_users" alias:"updated_by_user" on:"updated_by_user.id = core_translations.updated_by"`
	UpdatedAt      time.Time `json:"updated_at" sql:"updated_at"`
}

// CreateTranslationsFromStruct saves translations from struct to the database
func CreateTranslationsFromStruct(structureType, userID, languageCode string, object interface{}) error {
	objectType := reflect.TypeOf(object).Elem()
	objectValue := reflect.ValueOf(object).Elem()

	translations := []Translation{}
	for i := 0; i < objectType.NumField(); i++ {
		if objectType.Field(i).Tag.Get("table") == TableCoreTranslations {
			now := time.Now()
			structureID := objectValue.FieldByName("ID").Interface().(string)
			structureField := objectType.Field(i).Tag.Get("json")
			value := objectValue.Field(i).Interface().(string)
			translation := Translation{
				StructureID:    structureID,
				StructureField: structureField,
				StructureType:  structureType,
				Value:          value,
				LanguageCode:   languageCode,
				CreatedBy:      userID,
				UpdatedBy:      userID,
				CreatedAt:      now,
				UpdatedAt:      now,
			}
			translations = append(translations, translation)
		}
	}

	_, err := db.InsertStruct(TableCoreTranslations, translations)
	return err
}

// UpdateTranslationsFromStruct updates translations from struct to the database
func UpdateTranslationsFromStruct(structureType, userID, languageCode string, object interface{}, columns ...string) error {
	objectType := reflect.TypeOf(object).Elem()
	objectValue := reflect.ValueOf(object).Elem()

	for i := 0; i < objectType.NumField(); i++ {
		objectField := objectType.Field(i)
		if objectField.Tag.Get("table") == TableCoreTranslations {
			for _, column := range columns {
				if column == objectField.Tag.Get("json") {
					now := time.Now()
					structureID := objectValue.FieldByName("ID").Interface().(string)
					structureField := objectField.Tag.Get("json")
					value := objectValue.Field(i).Interface().(string)
					translation := Translation{
						StructureID:    structureID,
						StructureField: structureField,
						StructureType:  structureType,
						Value:          value,
						LanguageCode:   languageCode,
						UpdatedBy:      userID,
						UpdatedAt:      now,
					}

					structureIDColumn := fmt.Sprintf("%s.structure_id", TableCoreTranslations)
					structureFieldColumn := fmt.Sprintf("%s.structure_field", TableCoreTranslations)
					languageCodeColumn := fmt.Sprintf("%s.language_code", TableCoreTranslations)
					condition := builder.And(
						builder.Equal(structureIDColumn, structureID),
						builder.Equal(structureFieldColumn, structureField),
						builder.Equal(languageCodeColumn, languageCode),
					)

					err := db.UpdateStruct(
						TableCoreTranslations, &translation, condition,
						"structure_id", "structure_field", "structure_type", "value", "language_code",
					)
					if err != nil {
						return err
					}
					break
				}
			}
		}
	}

	return nil
}
