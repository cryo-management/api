package models

// Page defines the struct of this object
type Page struct {
	ID     string `json:"id" sql:"id" pk:"true"`
	ViewID string `json:"view_id" sql:"view_id" fk:"true"`
}
