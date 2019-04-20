package models

// Section defines the struct of this object
type Section struct {
	ID     string `json:"id" sql:"id" pk:"true"`
	ViewID string `json:"view_id" sql:"view_id" fk:"true"`
}

// SectionStructure defines the struct of this object
type SectionStructure struct {
	ID            string `json:"id" sql:"id" pk:"true"`
	SectionID     string `json:"section_id" sql:"section_id" fk:"true"`
	TabID         string `json:"tab_id" sql:"tab_id" fk:"true"`
	StructureID   string `json:"structure_id" sql:"structure_id" fk:"true"`
	StructureType string `json:"structure_type" sql:"structure_type"`
	Row           int    `json:"row" sql:"row"`
	Column        int    `json:"column" sql:"column"`
	Width         int    `json:"width" sql:"width"`
	Height        int    `json:"height" sql:"height"`
}
