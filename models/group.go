package models

type Group struct {
	ID          string `json:"id" sql:"id" pk:"true"`
	Name        string `json:"name" table:"translations" alias:"translations_name" sql:"value" on:"translations_name.structure_id = groups.id and translations_name.structure_field = 'name'"`
	Description string `json:"description" table:"translations" alias:"translations_description" sql:"value" on:"translations_description.structure_id = groups.id and translations_description.structure_field = 'description'"`
	Code        string `json:"code" sql:"code"`
	Active      bool   `json:"active" sql:"active"`
}

func (g *Group) GetID() string {
	return g.ID
}
