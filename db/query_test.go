package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type object struct {
	ID          string `json:"id" sql:"id" pk:"true"`
	Name        string `json:"name" type:"schemas" table:"translations" alias:"name" sql:"value" on:"name.structure_id = schema.id and name.structure_field = 'name'" external:"true" persist:"true"`
	Description string `json:"description" type:"schemas" table:"translations" alias:"description" sql:"value" on:"description.parent_id = schema.id and description.structure_field = 'description'" external:"true" persist:"true"`
	Code        string `json:"code" sql:"code"`
	Module      bool   `json:"module" sql:"module"`
	Active      bool   `json:"active" sql:"active"`
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
	obj := object{
		Code:   "C.001",
		Module: false,
		Active: true,
	}
	generatedQuery, _ := GenerateInsertQuery("table_name", obj)
	expectedQuery := "insert into table_name (code, module, active) values ($1, $2, $3)"
	assert.Equal(t, expectedQuery, generatedQuery, "invalid generated query")
}

func TestGenerateTranslationsInsertQuery(t *testing.T) {
	obj := object{
		ID:          "000001",
		Name:        "contract",
		Description: "text",
	}

	trs := translation{}

	generatedQuery, _ := GenerateTranslationsInsertQuery(obj.ID, "pt-br", obj, trs)
	expectedQuery := "insert into translations (structure_id, structure_type, structure_field, value, language_code) values ($1, $2, $3, $4, $5), ($6, $7, $8, $9, $10)"
	assert.Equal(t, expectedQuery, generatedQuery, "invalid generated query")
}
