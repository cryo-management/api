package models

import (
	"fmt"

	"github.com/cryo-management/api/common"
	"github.com/cryo-management/api/db"
)

// Group docs
type Group struct {
	ID          string `json:"id" sql:"id" pk:"true"`
	Name        string `json:"name" type:"groups" table:"translations" alias:"name" sql:"value" on:"translations_name.structure_id = groups.id and translations_name.structure_field = 'name'" external:"true" persist:"true"`
	Description string `json:"description" type:"groups" table:"translations" alias:"description" sql:"value" on:"translations_description.structure_id = groups.id and translations_description.structure_field = 'description'" external:"true" persist:"true"`
	Code        string `json:"code" sql:"code"`
	Active      bool   `json:"active" sql:"active"`
	Users       Users  `json:"users,omitempty"`
}

// Groups docs
type Groups []Group

// Create docs
func (g *Group) Create() error {
	table := "groups"
	query, args := db.GenerateInsertQuery(table, *g)
	conn := new(db.Database)
	id, err := conn.Insert(query, args...)
	if err != nil {
		return err
	}
	g.ID = id

	return nil
}

// Load docs
func (g *Group) Load(id string) error {
	table := "groups"
	sqlID := fmt.Sprintf("%s.id = '%s'", table, id)
	query := db.GenerateSelectQuery(table, Group{}, sqlID, fmt.Sprintf("and translations_name.language_code = '%s'", common.Session.User.Language), fmt.Sprintf("and translations_description.language_code = '%s'", common.Session.User.Language))
	conn := new(db.Database)
	rows, err := conn.Query(query)
	if err != nil {
		return err
	}

	err = db.StructScan(rows, g)
	if err != nil {
		return err
	}

	return nil
}

// Load docs
func (g *Groups) Load() error {
	table := "groups"
	query := db.GenerateSelectQuery(table, Group{}, fmt.Sprintf("translations_name.language_code = '%s'", common.Session.User.Language), fmt.Sprintf("and translations_description.language_code = '%s'", common.Session.User.Language))
	fmt.Printf(query)
	conn := new(db.Database)
	rows, err := conn.Query(query)
	if err != nil {
		return err
	}

	err = db.StructScan(rows, g)
	if err != nil {
		return err
	}

	return nil
}

// Delete docs
func (g *Group) Delete(id string) error {
	table := "groups"
	sqlID := fmt.Sprintf("%s.id = '%s'", table, id)
	query := db.GenerateDeleteQuery(table, sqlID)
	conn := new(db.Database)
	_, err := conn.Delete(query)
	if err != nil {
		return err
	}

	return nil
}
