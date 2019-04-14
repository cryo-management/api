package models

import (
	"fmt"

	"github.com/cryo-management/api/common"
	"github.com/cryo-management/api/db"
)

type Translation struct {
	ID             string `json:"id" sql:"id" pk:"true"`
	StructureID    string `json:"structure_id" sql:"structure_id" fk:"true"`
	StructureType  string `json:"structure_type" sql:"structure_type"`
	StructureField string `json:"structure_field" sql:"structure_field"`
	Value          string `json:"value" sql:"value"`
	LanguageCode   string `json:"language_code" sql:"language_code"`
}

func (t *Translation) Create(objID string, obj interface{}) error {
	query, args := db.GenerateTranslationsInsertQuery(objID, common.Session.User.Language, obj, Translation{})
	conn := new(db.Database)
	_, err := conn.Insert(query, args...)
	return err
}

func (t *Translation) DeleteByStructureID(structureID string) error {
	table := "translations"
	sqlID := fmt.Sprintf("%s.structure_id = '%s'", table, structureID)
	query := db.GenerateDeleteQuery(table, sqlID)
	conn := new(db.Database)
	_, err := conn.Delete(query)
	if err != nil {
		return err
	}

	return nil
}
