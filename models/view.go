package models

import (
	"fmt"

	"github.com/cryo-management/api/common"
	"github.com/cryo-management/api/db"
)

type View struct {
	ID               string `json:"id" sql:"id" pk:"true"`
	StructureID      string `json:"structure_id" sql:"structure_id" fk:"true"`
	StructureType    string `json:"structure_type" sql:"structure_type"`
	Name             string `json:"name" type:"views" table:"translations" alias:"name" sql:"value" on:"translations_name.structure_id = views.id and translations_name.structure_field = 'name'" external:"true" persist:"true"`
	Type             string `json:"type" sql:"type"`
	LastModifiedDate int    `json:"last_modified_date" sql:"last_modified_date"`
}

type Views []View

func (v *View) Create() error {
	table := "views"
	query, args := db.GenerateInsertQuery(table, *v)
	conn := new(db.Database)
	id, err := conn.Insert(query, args...)
	if err != nil {
		return err
	}
	v.ID = id

	return nil
}

func (v *View) Load(id string) error {
	table := "views"
	sqlID := fmt.Sprintf("%s.id = '%s'", table, id)
	query := db.GenerateSelectQuery(table, View{}, sqlID, fmt.Sprintf("and translations_name.language_code = '%s'", common.Session.User.Language), fmt.Sprintf("and translations_description.language_code = '%s'", common.Session.User.Language))
	conn := new(db.Database)
	rows, err := conn.Query(query)
	if err != nil {
		return err
	}

	err = db.StructScan(rows, v)
	if err != nil {
		return err
	}

	return nil
}

func (v *Views) Load(schemaID string) error {
	table := "views"
	sqlschemaID := fmt.Sprintf("%s.schema_id = '%s'", table, schemaID)
	query := db.GenerateSelectQuery(table, View{}, sqlschemaID, fmt.Sprintf("and translations_name.language_code = '%s'", common.Session.User.Language), fmt.Sprintf("and translations_description.language_code = '%s'", common.Session.User.Language))
	conn := new(db.Database)
	rows, err := conn.Query(query)
	if err != nil {
		return err
	}

	err = db.StructScan(rows, v)
	if err != nil {
		return err
	}

	return nil
}

func (v *View) Delete(id string) error {
	table := "views"
	sqlID := fmt.Sprintf("%s.id = '%s'", table, id)
	query := db.GenerateDeleteQuery(table, sqlID)
	conn := new(db.Database)
	_, err := conn.Delete(query)
	if err != nil {
		return err
	}

	return nil
}
