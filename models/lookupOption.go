package models

type LookupOption struct {
	ID       string `json:"id" sql:"id" pk:"true"`
	LookupID string `json:"lookup_id" sql:"lookup_id" fk:"true"`
	Value    string `json:"value" sql:"value"`
	Label    string `json:"label" table:"translations" alias:"translations_label" sql:"value" on:"translations_label.structure_id = lookups.id and translations_label.structure_field = 'label'"`
	Active   bool   `json:"active" sql:"active"`
}

func (l *LookupOption) GetID() string {
	return l.ID
}
