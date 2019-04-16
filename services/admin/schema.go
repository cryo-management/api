package admin

import (
	"github.com/cryo-management/api/models"
)

type SchemaService struct{}

func (s *SchemaService) Create(schema *models.Schema) error {
	err := schema.Create()
	if err != nil {
		return err
	}

	translationService := new(TranslationService)
	err = translationService.Create(*schema, schema.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *SchemaService) Load(schema *models.Schema, id string) error {
	err := schema.Load(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *SchemaService) LoadAll(schemas *models.Schemas) error {
	err := schemas.Load()
	if err != nil {
		return err
	}

	return nil
}

func (s *SchemaService) Delete(schema *models.Schema, id string) error {
	err := schema.Delete(id)
	if err != nil {
		return err
	}

	translationService := new(TranslationService)
	err = translationService.DeleteByStructureID(id)
	if err != nil {
		return err
	}

	groupPermission := new(models.GroupPermission)
	groupPermission.StructureID = id

	groupService := new(GroupService)
	err = groupService.DeletePermission(groupPermission)
	if err != nil {
		return err
	}

	fields := new(models.Fields)

	fieldService := new(FieldService)
	err = fieldService.DeleteBySchemaID(fields, id)
	if err != nil {
		return err
	}

	return nil
}
