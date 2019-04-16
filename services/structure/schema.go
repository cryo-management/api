package admin

import (
	"github.com/cryo-management/api/models"
)

type SchemaService struct{}

func (s *SchemaService) Load(schema *models.Schema, id string) error {
	err := schema.Load(id)
	if err != nil {
		return err
	}

	err = schema.Fields.LoadByPermission(schema.ID)
	if err != nil {
		return err
	}

	return nil
}
