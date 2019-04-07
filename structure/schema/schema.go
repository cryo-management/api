package schema

import (
	"github.com/cryo-management/api/db"
)

//Schema docs
type Schema struct {
	ID          string `json:"id" sql:"id" sqlType:"uuid" pk:"true"`
	Name        string `json:"name" type:"schemas" table:"translations" alias:"name" sql:"value" sqlType:"character varying" on:"name.structure_id = schemas.id and name.structure_field = 'name'" external:"true" persist:"true"`
	Description string `json:"description" type:"schemas" table:"translations" alias:"description" sql:"value" sqlType:"character varying" on:"description.parent_id = schemas.id and description.structure_field = 'description'" external:"true" persist:"true"`
	Code        string `json:"code" sql:"code" sqlType:"character varying"`
	Module      bool   `json:"module" sql:"module" sqlType:"boolean"`
	Active      bool   `json:"active" sql:"active" sqlType:"boolean"`
}

//Create docs
func (s *Schema) Create() (string, error) {
	query, args := db.GenerateInsertQuery("schemas", *s)
	conn := new(db.Database)
	id, err := conn.Insert(query, args...)
	if err != nil {
		return "", err
	}
	s.ID = id

	return id, nil
}

//GetAll docs
func (s *Schema) GetAll() ([]Schema, error) {
	query := db.GenerateSelectQuery("schemas", *s)
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
