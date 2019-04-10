package models

import (
	"fmt"

	"github.com/cryo-management/api/db"
)

// Schema docs
type Schema struct {
	ID               string  `json:"id" sql:"id" pk:"true"`
	Name             string  `json:"name" type:"schemas" table:"translations" alias:"name" sql:"value" on:"translations_name.structure_id = schemas.id and translations_name.structure_field = 'name'" external:"true" persist:"true"`
	Description      string  `json:"description" type:"schemas" table:"translations" alias:"description" sql:"value" on:"translations_description.structure_id = schemas.id and translations_description.structure_field = 'description'" external:"true" persist:"true"`
	Code             string  `json:"code" sql:"code"`
	Module           bool    `json:"module" sql:"module"`
	Active           bool    `json:"active" sql:"active"`
	LastModifiedDate int     `json:"last_modified_date" sql:"last_modified_date"`
	Fields           []Field `json:"fields"`
}

// Create docs
func (s *Schema) Create() (string, error) {
	table := "schemas"
	query, args := db.GenerateInsertQuery(table, *s)
	conn := new(db.Database)
	id, err := conn.Insert(query, args...)
	if err != nil {
		return "", err
	}
	s.ID = id

	return id, nil
}

// GetAll docs
func (s *Schema) GetAll() ([]Schema, error) {
	table := "schemas"
	query := db.GenerateSelectQuery(table, *s, "translations_name.language_code = 'pt-br'", "and translations_description.language_code = 'pt-br'")
	conn := new(db.Database)
	rows, err := conn.Query(query)
	if err != nil {
		return nil, err
	}

	schemaList := []Schema{}
	err = db.StructScan(rows, &schemaList)
	if err != nil {
		return nil, err
	}

	return schemaList, nil
}

// GetByCode docs
func (s *Schema) GetByCode(code string) (Schema, error) {
	table := "schemas"
	sqlCode := fmt.Sprintf("%s.code = '%s'", table, code)
	query := db.GenerateSelectQuery(table, *s, sqlCode, "and translations_name.language_code = 'pt-br'", "and translations_description.language_code = 'pt-br'")
	conn := new(db.Database)
	schema := Schema{}
	rows, err := conn.Query(query)
	if err != nil {
		return schema, err
	}

	err = db.StructScan(rows, &schema)
	if err != nil {
		return schema, err
	}

	return schema, nil
}

// Delete docs
func (s *Schema) Delete(code string) (Schema, error) {
	table := "schemas"
	sqlCode := fmt.Sprintf("%s.code = '%s'", table, code)
	query := db.GenerateSelectQuery(table, *s, sqlCode, "and translations_name.language_code = 'pt-br'", "and translations_description.language_code = 'pt-br'")
	conn := new(db.Database)
	schema := Schema{}
	rows, err := conn.Query(query)
	if err != nil {
		return schema, err
	}

	err = db.StructScan(rows, &schema)
	if err != nil {
		return schema, err
	}

	return schema, nil
}
