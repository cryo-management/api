package models

import (
	"reflect"

	"github.com/andreluzz/go-sql-builder/builder"
	"github.com/andreluzz/go-sql-builder/db"
)

// Translation defines the struct of this object
type Translation struct {
	ID             string `json:"id" sql:"id" pk:"true"`
	StructureID    string `json:"structure_id" sql:"structure_id" fk:"true"`
	StructureType  string `json:"structure_type" sql:"structure_type"`
	StructureField string `json:"structure_field" sql:"structure_field"`
	Value          string `json:"value" sql:"value"`
	LanguageCode   string `json:"language_code" sql:"language_code"`
}

// CreateTranslationsFromStruct saves translations from struct to the database
func CreateTranslationsFromStruct(structureType, languageCode string, object interface{}) error {
	objectType := reflect.TypeOf(object).Elem()
	objectValue := reflect.ValueOf(object).Elem()

	translations := []Translation{}
	for i := 0; i < objectType.NumField(); i++ {
		if objectType.Field(i).Tag.Get("table") == TableTranslations {
			structureID := objectValue.FieldByName("ID").Interface().(string)
			structureField := objectType.Field(i).Tag.Get("json")
			value := objectValue.Field(i).Interface().(string)
			translation := Translation{
				StructureID:    structureID,
				StructureField: structureField,
				StructureType:  structureType,
				Value:          value,
				LanguageCode:   languageCode,
			}
			translations = append(translations, translation)
		}
	}

	_, err := db.InsertStruct(TableTranslations, translations)
	return err
}

// UpdateTranslationsFromStruct updates translations from struct to the database
func UpdateTranslationsFromStruct(structureType, languageCode string, object interface{}, columns ...string) error {
	objectType := reflect.TypeOf(object).Elem()
	objectValue := reflect.ValueOf(object).Elem()

	for i := 0; i < objectType.NumField(); i++ {
		objectField := objectType.Field(i)
		if objectField.Tag.Get("table") == TableTranslations {
			for _, column := range columns {
				if column == objectField.Tag.Get("json") {
					structureID := objectValue.FieldByName("ID").Interface().(string)
					structureField := objectField.Tag.Get("json")
					value := objectValue.Field(i).Interface().(string)
					translation := Translation{
						StructureID:    structureID,
						StructureField: structureField,
						StructureType:  structureType,
						Value:          value,
						LanguageCode:   languageCode,
					}

					condition := builder.And(
						builder.Equal("translations.structure_id", structureID),
						builder.Equal("translations.structure_field", structureField),
					)

					err := db.UpdateStruct(
						TableTranslations, &translation, condition,
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
