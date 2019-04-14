package models

import (
	"fmt"

	"github.com/cryo-management/api/db"
)

type GroupPermission struct {
	ID             string `json:"id" sql:"id" pk:"true"`
	GroupID        string `json:"group_id" sql:"group_id" fk:"true"`
	StructureType  string `json:"structure_type" sql:"structure_type"`
	StructureID    string `json:"structure_id" sql:"structure_id" fk:"true"`
	Type           int    `json:"type" sql:"type"`
	ConditionQuery string `json:"condition_query" sql:"condition_query"`
}

func (g *GroupPermission) Create() error {
	table := "groups_permissions"
	query, args := db.GenerateInsertQuery(table, *g)
	conn := new(db.Database)
	id, err := conn.Insert(query, args...)
	if err != nil {
		return err
	}
	g.ID = id

	return nil
}

func (g *GroupPermission) Delete() error {
	table := "groups_permissions"
	sqlGroupID := fmt.Sprintf("%s.group_id = '%s'", table, g.GroupID)
	sqlStructureID := fmt.Sprintf("and %s.structure_id = '%s'", table, g.StructureID)
	sqlType := fmt.Sprintf("and %s.type = '%d'", table, g.Type)
	query := db.GenerateDeleteQuery(table, sqlGroupID, sqlStructureID, sqlType)
	conn := new(db.Database)
	_, err := conn.Delete(query)
	if err != nil {
		return err
	}

	return nil
}
