package models

import (
	"fmt"

	"github.com/cryo-management/api/db"
)

type User struct {
	ID        string `json:"id" sql:"id" pk:"true"`
	FirstName string `json:"first_name" sql:"first_name"`
	LastName  string `json:"last_name" sql:"last_name"`
	Email     string `json:"email" sql:"email"`
	Password  string `json:"password" sql:"password"`
	Language  string `json:"language" sql:"language"`
	Active    bool   `json:"active" sql:"active"`
	Groups    Groups `json:"groups,omitempty"`
}

type Users []User

func (u *User) Create() error {
	table := "users"
	query, args := db.GenerateInsertQuery(table, *u)
	conn := new(db.Database)
	id, err := conn.Insert(query, args...)
	if err != nil {
		return err
	}
	u.ID = id

	return nil
}

func (u *User) Load(id string) error {
	table := "users"
	sqlID := fmt.Sprintf("%s.id = '%s'", table, id)
	query := db.GenerateSelectQuery(table, *u, sqlID)
	conn := new(db.Database)
	rows, err := conn.Query(query)
	if err != nil {
		return err
	}

	err = db.StructScan(rows, u)
	if err != nil {
		return err
	}

	return nil
}

func (u *Users) Load() error {
	table := "users"
	query := db.GenerateSelectQuery(table, User{})
	conn := new(db.Database)
	rows, err := conn.Query(query)
	if err != nil {
		return err
	}

	err = db.StructScan(rows, u)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) Delete(id string) error {
	table := "users"
	sqlID := fmt.Sprintf("%s.id = '%s'", table, id)
	query := db.GenerateDeleteQuery(table, sqlID)
	conn := new(db.Database)
	_, err := conn.Delete(query)
	if err != nil {
		return err
	}

	return nil
}
