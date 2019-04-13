package models

import (
	"fmt"

	"github.com/cryo-management/api/db"
)

// GroupsUsers docs
type GroupsUsers struct {
	ID      string `json:"id" sql:"id" pk:"true"`
	GroupID string `json:"group_id" sql:"group_id" fk:"true"`
	UserID  string `json:"user_id" sql:"user_id" fk:"true"`
}

// AddUser docs
func (g *GroupsUsers) AddUser() error {
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

// RemoveUser docs
func (g *GroupsUsers) RemoveUser() error {
	table := "groups_users"
	sqlGroupID := fmt.Sprintf("%s.group_id = '%s'", table, g.GroupID)
	sqlUserID := fmt.Sprintf("and %s.user_id = '%s'", table, g.UserID)
	query := db.GenerateDeleteQuery(table, sqlGroupID, sqlUserID)
	fmt.Print(query)
	conn := new(db.Database)
	_, err := conn.Delete(query)
	if err != nil {
		return err
	}

	return nil
}
