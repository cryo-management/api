package models

// Lookup defines the struct of this object
type Lookup struct {
	ID           string `json:"id" sql:"id" pk:"true"`
	Name         string `json:"name" table:"translations" alias:"translations_name" sql:"value" on:"translations_name.structure_id = lookups.id and translations_name.structure_field = 'name'"`
	Description  string `json:"description" table:"translations" alias:"translations_description" sql:"value" on:"translations_description.structure_id = lookups.id and translations_description.structure_field = 'description'"`
	Code         string `json:"code" sql:"code"`
	Type         string `json:"type" sql:"type"`
	Query        string `json:"query" sql:"query"`
	Value        string `json:"value" sql:"value"`
	Label        string `json:"label" sql:"label"`
	Autocomplete string `json:"autocomplete" sql:"autocomplete"`
	Active       bool   `json:"active" sql:"active"`
}

// LookupOption defines the struct of this object
type LookupOption struct {
	ID       string `json:"id" sql:"id" pk:"true"`
	LookupID string `json:"lookup_id" sql:"lookup_id" fk:"true"`
	Value    string `json:"value" sql:"value"`
	Label    string `json:"label" table:"translations" alias:"translations_label" sql:"value" on:"translations_label.structure_id = lookups.id and translations_label.structure_field = 'label'"`
	Active   bool   `json:"active" sql:"active"`
}
