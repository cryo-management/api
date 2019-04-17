package models

//Schema defines the struct of this object
type Schema struct {
	ID          string `json:"id" sql:"id" pk:"true"`
	Name        string `json:"name" table:"translations" alias:"translations_name" sql:"value" on:"translations_name.structure_id = schemas.id and translations_name.structure_field = 'name'"`
	Description string `json:"description" table:"translations" alias:"translations_description" sql:"value" on:"translations_description.structure_id = schemas.id and translations_description.structure_field = 'description'"`
	Code        string `json:"code" sql:"code"`
	Module      bool   `json:"module" sql:"module"`
	Active      bool   `json:"active" sql:"active"`
}

//GetID returns object primary key
func (s *Schema) GetID() string {
	return s.ID
}
