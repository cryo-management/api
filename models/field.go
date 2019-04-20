package models

// Field defines the struct of this object
type Field struct {
	ID          string `json:"id" sql:"id" pk:"true"`
	SchemaID    string `json:"schema_id" sql:"schema_id" fk:"true"`
	Name        string `json:"name" table:"translations" alias:"translations_name" sql:"value" on:"translations_name.structure_id = fields.id and translations_name.structure_field = 'name'"`
	Description string `json:"description" table:"translations" alias:"translations_description" sql:"value" on:"translations_description.structure_id = fields.id and translations_description.structure_field = 'description'"`
	Code        string `json:"code" sql:"code"`
	FieldType   string `json:"field_type" sql:"field_type"`
	Multivalue  bool   `json:"multivalue" sql:"multivalue"`
	LookupID    string `json:"lookup_id" sql:"lookup_id" fk:"true"`
	Permission  int    `json:"permission"`
	Active      bool   `json:"active" sql:"active"`
}

