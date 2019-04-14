package models

import (
	"fmt"

	"github.com/cryo-management/api/common"
	"github.com/cryo-management/api/db"
)

type LookupOption struct {
	ID       string `json:"id" sql:"id" pk:"true"`
	LookupID string `json:"lookup_id" sql:"lookup_id" pk:"true"`
	Value    string `json:"value" sql:"value"`
	Label    string `json:"label" type:"lookups_options" table:"translations" alias:"label" sql:"value" on:"translations_label.structure_id = lookups.id and translations_label.structure_field = 'label'" external:"true" persist:"true"`
	Active   bool   `json:"active" sql:"active"`
}

type LookupOptions []LookupOption

func (l *LookupOption) Create() error {
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

func (l *LookupOption) Load(id string) error {
	table := "lookups"
	sqlID := fmt.Sprintf("%s.id = '%s'", table, id)
	query := db.GenerateSelectQuery(table, LookupOption{}, sqlID, fmt.Sprintf("and translations_name.language_code = '%s'", common.Session.User.Language), fmt.Sprintf("and translations_description.language_code = '%s'", common.Session.User.Language))
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

func (l *LookupOptions) Load() error {
	table := "lookups"
	query := db.GenerateSelectQuery(table, LookupOption{}, fmt.Sprintf("translations_name.language_code = '%s'", common.Session.User.Language), fmt.Sprintf("and translations_description.language_code = '%s'", common.Session.User.Language))
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

func (l *LookupOption) Delete(id string) error {
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
