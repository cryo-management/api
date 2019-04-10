package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type field struct {
	ID          string `json:"id" sql:"id" pk:"true"`
	SchemaID    string `json:"schema_id" sql:"schema_id" fk:"true"`
	Name        string `json:"name" type:"fields" table:"translations" alias:"name" sql:"value" on:"translations_name.structure_id = fields.id and translations_name.structure_field = 'name'" external:"true" persist:"true"`
	Description string `json:"description" type:"fields" table:"translations" alias:"description" sql:"value" on:"translations_description.structure_id = fields.id and translations_description.structure_field = 'description'" external:"true" persist:"true"`
	Code        string `json:"code" sql:"code"`
	Type        string `json:"type" sql:"type"`
	Multivalue  string `json:"multivalue" sql:"multivalue"`
	LookupID    string `json:"lookup_id" sql:"lookup_id"`
	Permission  string `json:"permission" sql:"permission"`
	Active      string `json:"active" sql:"active"`
}

type schemaComplete struct {
	ID               string  `json:"id" sql:"id" pk:"true"`
	Name             string  `json:"name" type:"schemas" table:"translations" alias:"name" sql:"value" on:"translations_name.structure_id = schemas.id and translations_name.structure_field = 'name'" external:"true" persist:"true"`
	Description      string  `json:"description" type:"schemas" table:"translations" alias:"description" sql:"value" on:"translations_description.structure_id = schemas.id and translations_description.structure_field = 'description'" external:"true" persist:"true"`
	Code             string  `json:"code" sql:"code"`
	Module           bool    `json:"module" sql:"module"`
	Active           bool    `json:"active" sql:"active"`
	LastModifiedDate int     `json:"last_modified_date" sql:"last_modified_date"`
	Fields           []field `json:"fields"`
}

type schemaSimple struct {
	ID               string  `json:"id" sql:"id" pk:"true"`
	Code             string  `json:"code" sql:"code"`
	Module           bool    `json:"module" sql:"module"`
	Active           bool    `json:"active" sql:"active"`
	LastModifiedDate int     `json:"last_modified_date" sql:"last_modified_date"`
	Fields           []field `json:"fields"`
}

type translation struct {
	ID             string `json:"id" sql:"id" pk:"true"`
	StructureID    string `json:"structure_id" sql:"structure_id"`
	StructureType  string `json:"structure_type" sql:"structure_type"`
	StructureField string `json:"structure_field" sql:"structure_field"`
	Value          string `json:"value" sql:"value"`
	LanguageCode   string `json:"language_code" sql:"language_code"`
}

func TestGenerateInsertQuery(t *testing.T) {
	obj := schemaComplete{
		Code:   "C.001",
		Module: false,
		Active: true,
	}
	generatedQuery, _ := GenerateInsertQuery("table_name", obj)
	expectedQuery := "insert into table_name (code, module, active) values ($1, $2, $3)"
	assert.Equal(t, expectedQuery, generatedQuery, "invalid generated query")
}

func TestGenerateSelectQueryComplete(t *testing.T) {
	obj := schemaComplete{}
	generatedQuery := GenerateSelectQuery("schemas", obj, "translations_name.language_code = 'pt-br'", "and translations_description.language_code = 'pt-br'")
	expectedQuery := "select schemas.id, translations_name.value as name, translations_description.value as description, schemas.code, schemas.module, schemas.active, schemas.last_modified_date from schemas join translations translations_name on translations_name.structure_id = schemas.id and translations_name.structure_field = 'name' join translations translations_description on translations_description.structure_id = schemas.id and translations_description.structure_field = 'description' where translations_name.language_code = 'pt-br' and translations_description.language_code = 'pt-br'"
	assert.Equal(t, expectedQuery, generatedQuery, "invalid generated query")
}

func TestGenerateSelectQuerySimple(t *testing.T) {
	obj := schemaSimple{}
	generatedQuery := GenerateSelectQuery("schemas", obj)
	expectedQuery := "select schemas.id, schemas.code, schemas.module, schemas.active, schemas.last_modified_date from schemas"
	assert.Equal(t, expectedQuery, generatedQuery, "invalid generated query")
}

func TestGenerateTranslationsInsertQuery(t *testing.T) {
	obj := schemaComplete{
		ID:          "000001",
		Name:        "contract",
		Description: "text",
	}

	trs := translation{}

	generatedQuery, _ := GenerateTranslationsInsertQuery(obj.ID, "pt-br", obj, trs)
	expectedQuery := "insert into translations (structure_id, structure_type, structure_field, value, language_code) values ($1, $2, $3, $4, $5), ($6, $7, $8, $9, $10)"
	assert.Equal(t, expectedQuery, generatedQuery, "invalid generated query")
}
