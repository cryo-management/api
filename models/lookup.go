package models

import (
	"fmt"

	"github.com/cryo-management/api/common"
	"github.com/cryo-management/api/db"
)

type Lookup struct {
	ID           string       `json:"id" sql:"id" pk:"true"`
	Name         string       `json:"name" type:"lookups" table:"translations" alias:"name" sql:"value" on:"translations_name.structure_id = lookups.id and translations_name.structure_field = 'name'" external:"true" persist:"true"`
	Description  string       `json:"description" type:"lookups" table:"translations" alias:"description" sql:"value" on:"translations_description.structure_id = lookups.id and translations_description.structure_field = 'description'" external:"true" persist:"true"`
	Code         string       `json:"code" sql:"code"`
	Type         string       `json:"type" sql:"type"`
	Query        string       `json:"query" sql:"query"`
	Value        string       `json:"value" sql:"value"`
	Label        string       `json:"label" sql:"label"`
	Autocomplete string       `json:"autocomplete" sql:"autocomplete"`
	Active       bool         `json:"active" sql:"active"`
	LookupOption LookupOption `json:"lookups_options,omitempty"`
}

type Lookups []Lookup

func (l *Lookup) Create() error {
	table := "lookups"
	query, args := db.GenerateInsertQuery(table, *l)
	conn := new(db.Database)
	id, err := conn.Insert(query, args...)
	if err != nil {
		return err
	}
	l.ID = id

	return nil
}

func (l *Lookup) Load(id string) error {
	table := "lookups"
	sqlID := fmt.Sprintf("%s.id = '%s'", table, id)
	query := db.GenerateSelectQuery(table, Lookup{}, sqlID, fmt.Sprintf("and translations_name.language_code = '%s'", common.Session.User.Language), fmt.Sprintf("and translations_description.language_code = '%s'", common.Session.User.Language))
	conn := new(db.Database)
	rows, err := conn.Query(query)
	if err != nil {
		return err
	}

	err = db.StructScan(rows, l)
	if err != nil {
		return err
	}

	return nil
}

func (l *Lookups) Load() error {
	table := "lookups"
	query := db.GenerateSelectQuery(table, Lookup{}, fmt.Sprintf("translations_name.language_code = '%s'", common.Session.User.Language), fmt.Sprintf("and translations_description.language_code = '%s'", common.Session.User.Language))
	conn := new(db.Database)
	rows, err := conn.Query(query)
	if err != nil {
		return err
	}

	err = db.StructScan(rows, l)
	if err != nil {
		return err
	}

	return nil
}

func (l *Lookup) Delete(id string) error {
	table := "lookups"
	sqlID := fmt.Sprintf("%s.id = '%s'", table, id)
	query := db.GenerateDeleteQuery(table, sqlID)
	conn := new(db.Database)
	_, err := conn.Delete(query)
	if err != nil {
		return err
	}

	return nil
}
