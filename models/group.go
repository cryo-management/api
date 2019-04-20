package models

// Group defines the struct of this object
type Group struct {
	ID          string `json:"id" sql:"id" pk:"true"`
	Name        string `json:"name" table:"translations" alias:"translations_name" sql:"value" on:"translations_name.structure_id = groups.id and translations_name.structure_field = 'name'"`
	Description string `json:"description" table:"translations" alias:"translations_description" sql:"value" on:"translations_description.structure_id = groups.id and translations_description.structure_field = 'description'"`
	Code        string `json:"code" sql:"code"`
	Active      bool   `json:"active" sql:"active"`
}

// Permission defines the struct of this object
type Permission struct {
	ID             string `json:"id" sql:"id" pk:"true"`
	GroupID        string `json:"group_id" sql:"group_id" fk:"true"`
	StructureType  string `json:"structure_type" sql:"structure_type"`
	StructureID    string `json:"structure_id" sql:"structure_id" fk:"true"`
	Type           int    `json:"type" sql:"type"`
	ConditionQuery string `json:"condition_query" sql:"condition_query"`
}

