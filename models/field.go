package models

//Field docs
type Field struct {
	ID          string `json:"id" sql:"id" sqlType:"uuid" pk:"true"`
	SchemaID    string `json:"schema_id" sql:"schema_id" sqlType:"uuid" fk:"true"`
	Name        string `json:"name" type:"schemas" table:"translations" alias:"name" sql:"value" sqlType:"character varying" on:"name.structure_id = schemas.id and name.structure_field = 'name'" external:"true" persist:"true"`
	Description string `json:"description" type:"schemas" table:"translations" alias:"description" sql:"value" sqlType:"character varying" on:"description.parent_id = schemas.id and description.structure_field = 'description'" external:"true" persist:"true"`
	Code        string `json:"code" sql:"code" sqlType:"character varying"`
}
