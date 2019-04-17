package models

import "time"

//View defines the struct of this object
type View struct {
	ID               string    `json:"id" sql:"id" pk:"true"`
	StructureID      string    `json:"structure_id" sql:"structure_id" fk:"true"`
	StructureType    string    `json:"structure_type" sql:"structure_type"`
	Name             string    `json:"name" table:"translations" alias:"translations_name" sql:"value" on:"translations_name.structure_id = views.id and translations_name.structure_field = 'name'"`
	Type             string    `json:"type" sql:"type"`
	LastModifiedDate time.Time `json:"last_modified_date" sql:"last_modified_date"`
}

//GetID returns object primary key
func (v *View) GetID() string {
	return v.ID
}
