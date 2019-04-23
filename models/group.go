package models

import "time"

// Group defines the struct of this object
type Group struct {
	ID            string    `json:"id" sql:"id" pk:"true"`
	Code          string    `json:"code" sql:"code"`
	Name          string    `json:"name" table:"core_translations" alias:"core_translations_name" sql:"value" on:"core_translations_name.structure_id = core_groups.id and core_translations_name.structure_field = 'name'"`
	Description   string    `json:"description" table:"core_translations" alias:"core_translations_description" sql:"value" on:"core_translations_description.structure_id = core_groups.id and core_translations_description.structure_field = 'description'"`
	Active        bool      `json:"active" sql:"active"`
	CreatedByUser *User     `json:"created_by_user" table:"core_groups" alias:"created_by_user" on:"created_by_user.id = core_groups.created_by"`
	CreatedAt     time.Time `json:"created_at" sql:"created_at"`
	UpdatedByUser *User     `json:"updated_by_user" table:"core_groups" alias:"updated_by_user" on:"updated_by_user.id = core_groups.updated_by"`
	UpdatedAt     time.Time `json:"updated_at" sql:"updated_at"`
}

// Permission defines the struct of this object
type Permission struct {
	ID             string    `json:"id" sql:"id" pk:"true"`
	GroupID        string    `json:"group_id" sql:"group_id" fk:"true"`
	StructureType  string    `json:"structure_type" sql:"structure_type"`
	StructureID    string    `json:"structure_id" sql:"structure_id" fk:"true"`
	Type           int       `json:"type" sql:"type"`
	ConditionQuery string    `json:"condition_query" sql:"condition_query"`
	CreatedByUser  *User     `json:"created_by_user" table:"core_grp_permissions" alias:"created_by_user" on:"created_by_user.id = core_grp_permissions.created_by"`
	CreatedAt      time.Time `json:"created_at" sql:"created_at"`
	UpdatedByUser  *User     `json:"updated_by_user" table:"core_grp_permissions" alias:"updated_by_user" on:"updated_by_user.id = core_grp_permissions.updated_by"`
	UpdatedAt      time.Time `json:"updated_at" sql:"updated_at"`
}
