package models

import "time"

// Page defines the struct of this object
type Page struct {
	ID            string    `json:"id" sql:"id" pk:"true"`
	Code          string    `json:"code" sql:"code"`
	Name          string    `json:"name" table:"core_translations" alias:"core_translations_name" sql:"value" on:"core_translations_name.structure_id = core_sch_pages.id and core_translations_name.structure_field = 'name'"`
	SchemaID      string    `json:"schema_id" sql:"schema_id" fk:"true"`
	Description   string    `json:"description" table:"core_translations" alias:"core_translations_description" sql:"value" on:"core_translations_description.structure_id = core_sch_pages.id and core_translations_description.structure_field = 'description'"`
	Type          string    `json:"type" sql:"type"`
	Active        bool      `json:"active" sql:"active"`
	CreatedByUser *User     `json:"created_by_user" table:"core_sch_pages" alias:"created_by_user" on:"created_by_user.id = core_sch_pages.created_by"`
	CreatedAt     time.Time `json:"created_at" sql:"created_at"`
	UpdatedByUser *User     `json:"updated_by_user" table:"core_sch_pages" alias:"updated_by_user" on:"updated_by_user.id = core_sch_pages.updated_by"`
	UpdatedAt     time.Time `json:"updated_at" sql:"updated_at"`
}

// Section defines the struct of this object
type Section struct {
	ID            string    `json:"id" sql:"id" pk:"true"`
	Code          string    `json:"code" sql:"code"`
	Name          string    `json:"name" table:"core_translations" alias:"core_translations_name" sql:"value" on:"core_translations_name.structure_id = core_sch_pag_sections.id and core_translations_name.structure_field = 'name'"`
	Description   string    `json:"description" table:"core_translations" alias:"core_translations_description" sql:"value" on:"core_translations_description.structure_id = core_sch_pag_sections.id and core_translations_description.structure_field = 'description'"`
	SchemaID      string    `json:"schema_id" sql:"schema_id" fk:"true"`
	PageID        string    `json:"core_page_id" sql:"core_page_id" fk:"true"`
	CreatedByUser *User     `json:"created_by_user" table:"core_sch_pag_sections" alias:"created_by_user" on:"created_by_user.id = core_sch_pag_sections.created_by"`
	CreatedAt     time.Time `json:"created_at" sql:"created_at"`
	UpdatedByUser *User     `json:"updated_by_user" table:"core_sch_pag_sections" alias:"updated_by_user" on:"updated_by_user.id = core_sch_pag_sections.updated_by"`
	UpdatedAt     time.Time `json:"updated_at" sql:"updated_at"`
}

// Tab defines the struct of this object
type Tab struct {
	ID            string    `json:"id" sql:"id" pk:"true"`
	Code          string    `json:"code" sql:"code"`
	Name          string    `json:"name" table:"core_translations" alias:"core_translations_name" sql:"value" on:"core_translations_name.structure_id = core_sch_pag_sec_tabs.id and core_translations_name.structure_field = 'name'"`
	Description   string    `json:"description" table:"core_translations" alias:"core_translations_description" sql:"value" on:"core_translations_description.structure_id = core_sch_pag_sec_tabs.id and core_translations_description.structure_field = 'description'"`
	SchemaID      string    `json:"schema_id" sql:"schema_id" fk:"true"`
	PageID        string    `json:"core_page_id" sql:"core_page_id" fk:"true"`
	CreatedByUser *User     `json:"created_by_user" table:"core_sch_pag_sec_tabs" alias:"created_by_user" on:"created_by_user.id = core_sch_pag_sec_tabs.created_by"`
	CreatedAt     time.Time `json:"created_at" sql:"created_at"`
	UpdatedByUser *User     `json:"updated_by_user" table:"core_sch_pag_sec_tabs" alias:"updated_by_user" on:"updated_by_user.id = core_sch_pag_sec_tabs.updated_by"`
	UpdatedAt     time.Time `json:"updated_at" sql:"updated_at"`
}

// SectionStructure defines the struct of this object
type SectionStructure struct {
	ID            string    `json:"id" sql:"id" pk:"true"`
	PageID        string    `json:"core_page_id" sql:"core_page_id" fk:"true"`
	SchemaID      string    `json:"schema_id" sql:"schema_id" fk:"true"`
	SectionID     string    `json:"core_section_id" sql:"core_section_id" fk:"true"`
	TabID         string    `json:"core_tab_id" sql:"core_tab_id" fk:"true"`
	StructureID   string    `json:"structure_id" sql:"structure_id" fk:"true"`
	StructureType string    `json:"structure_type" sql:"structure_type"`
	Row           int       `json:"row" sql:"row"`
	Column        int       `json:"column" sql:"column"`
	Width         int       `json:"width" sql:"width"`
	Height        int       `json:"height" sql:"height"`
	CreatedByUser *User     `json:"created_by_user" table:"core_sch_pag_sec_structures" alias:"created_by_user" on:"created_by_user.id = core_sch_pag_sec_structures.created_by"`
	CreatedAt     time.Time `json:"created_at" sql:"created_at"`
	UpdatedByUser *User     `json:"updated_by_user" table:"core_sch_pag_sec_structures" alias:"updated_by_user" on:"updated_by_user.id = core_sch_pag_sec_structures.updated_by"`
	UpdatedAt     time.Time `json:"updated_at" sql:"updated_at"`
}
