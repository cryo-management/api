package schema

import (
	"github.com/cryo-management/api/db"
)

//Schema docs
type Schema struct {
	ID          string `json:"id" sql:"id" editable:"false"`
	Name        string `json:"name" table:"translations" alias:"name" sql:"value" on:"name.structure_id = schema.id and name.structure_field = 'name'" embedded:"true" persist:"true"`
	Description string `json:"description" table:"translations" alias:"description" sql:"value" on:"description.parent_id = schema.id and description.structure_field = 'description'" embedded:"true" persist:"true"`
	Code        string `json:"code" sql:"code"`
	Module      bool   `json:"module" sql:"module"`
	Active      bool   `json:"active" sql:"active"`
}

//Create docs
func (s *Schema) Create() (string, error) {
	query, args := db.GenerateInsertQuery(s)
	conn := new(db.Database)
	id, err := conn.Insert(query, args...)
	if err != nil {
		return "", err
	}
	s.ID = id

	return id, nil
}
