package translation

import (
	"github.com/cryo-management/api/db"
)

//Translation docs
type Translation struct {
	ID             string `json:"id" sql:"id" pk:"true"`
	StructureID    string `json:"structure_id" sql:"structure_id"`
	StructureType  string `json:"structure_type" sql:"structure_type"`
	StructureField string `json:"structure_field" sql:"structure_field"`
	Value          string `json:"value" sql:"value"`
	LanguageCode   string `json:"language_code" sql:"language_code"`
}

//Save docs
func Save(objID, langCode string, obj interface{}) error {
	query, args := db.GenerateTranslationsInsertQuery(objID, langCode, obj, Translation{})
	conn := new(db.Database)
	_, err := conn.Insert(query, args...)
	return err
}
