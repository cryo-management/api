package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type schema struct {
	ID          string `json:"id" sql:"id" sqlType:"uuid" pk:"true"`
	Name        string `json:"name" type:"schemas" table:"translations" alias:"name" sql:"value" sqlType:"character varying" on:"name.structure_id = schemas.id and name.structure_field = 'name'" external:"true" persist:"true"`
	Description string `json:"description" type:"schemas" table:"translations" alias:"description" sql:"value" sqlType:"character varying" on:"description.parent_id = schemas.id and description.structure_field = 'description'" external:"true" persist:"true"`
	Code        string `json:"code" sql:"code" sqlType:"character varying"`
	Module      bool   `json:"module" sql:"module" sqlType:"boolean"`
	Active      bool   `json:"active" sql:"active" sqlType:"boolean"`
}

type translation struct {
	ID             string `json:"id" sql:"id" sqlType:"uuid" pk:"true"`
	StructureID    string `json:"structure_id" sql:"structure_id" sqlType:"uuid"`
	StructureType  string `json:"structure_type" sql:"structure_type" sqlType:"character varying"`
	StructureField string `json:"structure_field" sql:"structure_field" sqlType:"character varying"`
	Value          string `json:"value" sql:"value" sqlType:"character varying"`
	LanguageCode   string `json:"language_code" sql:"language_code" sqlType:"character varying"`
}

func TestGenerateInsertQuery(t *testing.T) {
	obj := schema{
		Code:   "C.001",
		Module: false,
		Active: true,
	}
	generatedQuery, _ := GenerateInsertQuery("table_name", obj)
	expectedQuery := "insert into table_name (code, module, active) values ($1, $2, $3)"
	assert.Equal(t, expectedQuery, generatedQuery, "invalid generated query")
}

func TestGenerateSelectQuery(t *testing.T) {
	obj := schema{}
	generatedQuery := GenerateSelectQuery("table_name", obj)
	expectedQuery := "select * from crosstab($$select table_name.id, table_name.code, table_name.module, table_name.active, translations.structure_field, translations.value from table_name join translations on translations.structure_id = table_name.id and translations.structure_field in ('name', 'description')$$, $$values ('name'), ('description')$$) as tab (id uuid, code character varying, module boolean, active boolean, name character varying, description character varying)"
	assert.Equal(t, expectedQuery, generatedQuery, "invalid generated query")
}

func TestGenerateTranslationsInsertQuery(t *testing.T) {
	obj := schema{
		ID:          "000001",
		Name:        "contract",
		Description: "text",
	}

	trs := translation{}

	generatedQuery, _ := GenerateTranslationsInsertQuery(obj.ID, "pt-br", obj, trs)
	expectedQuery := "insert into translations (structure_id, structure_type, structure_field, value, language_code) values ($1, $2, $3, $4, $5), ($6, $7, $8, $9, $10)"
	assert.Equal(t, expectedQuery, generatedQuery, "invalid generated query")
}
