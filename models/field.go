package models

import (
	"fmt"

	"github.com/cryo-management/api/db"
)

//Field docs
type Field struct {
	ID          string `json:"id" sql:"id" pk:"true"`
	SchemaID    string `json:"schema_id" sql:"schema_id" fk:"true"`
	Name        string `json:"name" type:"fields" table:"translations" alias:"name" sql:"value" on:"translations_name.structure_id = fields.id and translations_name.structure_field = 'name'" external:"true" persist:"true"`
	Description string `json:"description" type:"fields" table:"translations" alias:"description" sql:"value" on:"translations_description.structure_id = fields.id and translations_description.structure_field = 'description'" external:"true" persist:"true"`
	Code        string `json:"code" sql:"code"`
	Type        string `json:"type" sql:"type"`
	Multivalue  bool   `json:"multivalue" sql:"multivalue"`
	LookupID    string `json:"lookup_id" sql:"lookup_id"`
	Permission  int    `json:"permission" sql:"permission"`
	Active      bool   `json:"active" sql:"active"`
}

// Create docs
func (field *Field) Create() (string, error) {
	table := "fields"
	query, args := db.GenerateInsertQuery(table, *field)
	conn := new(db.Database)
	id, err := conn.Insert(query, args...)
	if err != nil {
		return "", err
	}
	field.ID = id

	return id, nil
}

// GetAllBySchemaID docs
func (field *Field) GetAllBySchemaID(schemaID string, fieldList *[]Field) error {
	table := "fields"
	sqlschemaID := fmt.Sprintf("%s.schema_id = '%s'", table, schemaID)
	query := db.GenerateSelectQuery(table, *field, sqlschemaID, "and translations_name.language_code = 'pt-br'", "and translations_description.language_code = 'pt-br'")
	conn := new(db.Database)
	rows, err := conn.Query(query)
	if err != nil {
		return err
	}

	err = db.StructScan(rows, &fieldList)
	if err != nil {
		return err
	}

	return nil
}

// GetAllWithPermissionBySchemaID docs
func (field *Field) GetAllWithPermissionBySchemaID(schemaID string, fieldList *[]Field) error {
	query := "select f.id, f.schema_id, f.code, translations_name.value as name, translations_description.value as description, f.type, f.multivalue, f.lookup_id, f.active, max(gp.type) permission from user_group ug join groups g on g.id = ug.group_id join groups_permissions gp on gp.group_id = ug.group_id join fields f on f.id = gp.structure_id and gp.structure_type = 'field' join schemas s on s.id = f.schema_id join translations translations_name on translations_name.structure_id = f.id and translations_name.structure_field = 'name' join translations translations_description on translations_description.structure_id = f.id and translations_description.structure_field = 'description' where f.active = true and g.active = true and f.schema_id = $1 and ug.user_id = $2 and translations_name.language_code = $3 and translations_description.language_code = $4 group by f.id, f.schema_id, f.code, translations_name.value, translations_description.value, f.type, f.multivalue, f.lookup_id, f.active"
	conn := new(db.Database)
	rows, err := conn.Query(query, schemaID, "059fa339-025c-4104-ab55-c764d3028f63", "pt-br", "pt-br")
	if err != nil {
		return err
	}

	err = db.StructScan(rows, &fieldList)
	if err != nil {
		return err
	}

	return nil
}
