package models

import (
	"fmt"

	"github.com/cryo-management/api/common"
	"github.com/cryo-management/api/db"
)

type Schema struct {
	ID               string `json:"id" sql:"id" pk:"true"`
	Name             string `json:"name" type:"schemas" table:"translations" alias:"name" sql:"value" on:"translations_name.structure_id = schemas.id and translations_name.structure_field = 'name'" external:"true" persist:"true"`
	Description      string `json:"description" type:"schemas" table:"translations" alias:"description" sql:"value" on:"translations_description.structure_id = schemas.id and translations_description.structure_field = 'description'" external:"true" persist:"true"`
	Code             string `json:"code" sql:"code"`
	Module           bool   `json:"module" sql:"module"`
	Active           bool   `json:"active" sql:"active"`
	LastModifiedDate int    `json:"last_modified_date" sql:"last_modified_date" readOnly:"true"`
	Fields           Fields `json:"fields,omitempty"`
}

type Schemas []Schema

func (s *Schema) Create() error {
	table := "schemas"
	query, args := db.GenerateInsertQuery(table, *s)
	conn := new(db.Database)
	id, err := conn.Insert(query, args...)
	if err != nil {
		return err
	}
	s.ID = id

	return nil
}

func (s *Schema) Load(id string) error {
	table := "schemas"
	sqlID := fmt.Sprintf("%s.id = '%s'", table, id)
	query := db.GenerateSelectQuery(table, *s, sqlID, fmt.Sprintf("and translations_name.language_code = '%s'", common.Session.User.Language), fmt.Sprintf("and translations_description.language_code = '%s'", common.Session.User.Language))
	conn := new(db.Database)
	rows, err := conn.Query(query)
	if err != nil {
		return err
	}

	err = db.StructScan(rows, s)
	if err != nil {
		return err
	}

	return nil
}

func (s *Schemas) Load() error {
	table := "schemas"
	query := db.GenerateSelectQuery(table, Schema{}, fmt.Sprintf("translations_name.language_code = '%s'", common.Session.User.Language), fmt.Sprintf("and translations_description.language_code = '%s'", common.Session.User.Language))
	conn := new(db.Database)
	rows, err := conn.Query(query)
	if err != nil {
		return err
	}

	err = db.StructScan(rows, s)
	if err != nil {
		return err
	}

	return nil
}

func (s *Schema) Delete(id string) error {
	table := "schemas"
	sqlID := fmt.Sprintf("%s.id = '%s'", table, id)
	query := db.GenerateDeleteQuery(table, sqlID)
	conn := new(db.Database)
	_, err := conn.Delete(query)
	if err != nil {
		return err
	}

	return nil
}
