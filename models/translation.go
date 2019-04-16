package models

import (
	"reflect"

	"github.com/andreluzz/go-sql-builder/db"
)

type Translation struct {
	ID             string `json:"id" sql:"id" pk:"true"`
	StructureID    string `json:"structure_id" sql:"structure_id" fk:"true"`
	StructureType  string `json:"structure_type" sql:"structure_type"`
	StructureField string `json:"structure_field" sql:"structure_field"`
	Value          string `json:"value" sql:"value"`
	LanguageCode   string `json:"language_code" sql:"language_code"`
}

func CreateTranslationsFromStruct(structureType, languageCode string, model Model) error {
	modelType := reflect.TypeOf(model).Elem()
	modelValue := reflect.ValueOf(model).Elem()

	translations := []Translation{}
	for i := 0; i < modelType.NumField(); i++ {
		if modelType.Field(i).Tag.Get("table") == TableTranslations {
			trs := Translation{
				StructureID:    modelValue.FieldByName("ID").Interface().(string),
				StructureField: modelType.Field(i).Tag.Get("json"),
				StructureType:  structureType,
				Value:          modelValue.Field(i).Interface().(string),
				LanguageCode:   languageCode,
			}
			translations = append(translations, trs)
		}
	}

	_, err := db.InsertStruct(TableTranslations, translations)
	return err
}

func (t *Translation) DeleteByStructureID(structureID string) error {
	return nil
}
