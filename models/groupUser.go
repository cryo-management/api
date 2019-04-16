package models

import (
	"fmt"

	"github.com/cryo-management/api/db"
)

type GroupUser struct {
	ID      string `json:"id" sql:"id" pk:"true"`
	GroupID string `json:"group_id" sql:"group_id" fk:"true"`
	UserID  string `json:"user_id" sql:"user_id" fk:"true"`
}

func (g *GroupUser) Create() error {
	table := "groups_users"
	query, args := db.GenerateInsertQuery(table, *g)
	conn := new(db.Database)
	id, err := conn.Insert(query, args...)
	if err != nil {
		return err
	}
	g.ID = id

	return nil
}

func (g *GroupUser) Delete() error {
	table := "groups_users"
	sqlGroupID, sqlUserID := "", ""
	if g.GroupID != "" && g.UserID != "" {
		sqlGroupID = fmt.Sprintf("%s.group_id = '%s'", table, g.GroupID)
		sqlUserID = fmt.Sprintf("and %s.user_id = '%s'", table, g.UserID)
	} else {
		if g.GroupID != "" {
			sqlGroupID = fmt.Sprintf("%s.group_id = '%s'", table, g.GroupID)
		}
		if g.UserID != "" {
			sqlUserID = fmt.Sprintf("%s.user_id = '%s'", table, g.UserID)
		}
	}
	query := db.GenerateDeleteQuery(table, sqlGroupID, sqlUserID)
	conn := new(db.Database)
	_, err := conn.Delete(query)
	if err != nil {
		return err
	}

	return nil
}
