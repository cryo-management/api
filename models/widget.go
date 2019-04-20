package models

// Widget defines the struct of this object
type Widget struct {
	ID   string `json:"id" sql:"id" pk:"true"`
	Type bool   `json:"type" sql:"type"`
}
