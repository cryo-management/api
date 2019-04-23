package models

import "time"

// Widget defines the struct of this object
type Widget struct {
	ID            string    `json:"id" sql:"id" pk:"true"`
	Code          string    `json:"code" sql:"code"`
	Type          string    `json:"type" sql:"type"`
	Active        bool      `json:"active" sql:"active"`
	CreatedByUser *User     `json:"created_by_user" table:"core_widgets" alias:"created_by_user" on:"created_by_user.id = core_widgets.created_by"`
	CreatedAt     time.Time `json:"created_at" sql:"created_at"`
	UpdatedByUser *User     `json:"updated_by_user" table:"core_widgets" alias:"updated_by_user" on:"updated_by_user.id = core_widgets.updated_by"`
	UpdatedAt     time.Time `json:"updated_at" sql:"updated_at"`
}
