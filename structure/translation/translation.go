package translation

import (
	"github.com/cryo-management/api/db"
)

//Translation docs
type Translation struct {
	ID             string
	StructureID    string
	StructureField string
}

//Save docs
func Save(obj interface{}) error {
	query, args := db.GenerateTranslationsInsertQuery(obj)
	conn := new(db.Database)
	_, err := conn.Insert(query, args...)
	return err
}
