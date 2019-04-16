package models

import "github.com/andreluzz/go-sql-builder/db"

type Schema struct {
	ID          string `json:"id" sql:"id" pk:"true"`
	Name        string `json:"name" type:"schemas" table:"translations" alias:"name" sql:"value" on:"translations_name.structure_id = schemas.id and translations_name.structure_field = 'name'" external:"true" persist:"true"`
	Description string `json:"description" type:"schemas" table:"translations" alias:"description" sql:"value" on:"translations_description.structure_id = schemas.id and translations_description.structure_field = 'description'" external:"true" persist:"true"`
	Code        string `json:"code" sql:"code"`
	Module      bool   `json:"module" sql:"module"`
	Active      bool   `json:"active" sql:"active"`
}

type Schemas []Schema

func (s *Schema) GetID() string {
	return s.ID
}

func (s *Schema) Create() (string, error) {
	return db.InsertStruct(TableSchema, s)
}

func (s *Schema) Load(id string) error {
	return nil
}

func (s *Schemas) Load() error {
	return nil
}

func (s *Schema) Delete(id string) error {
	return nil
}
