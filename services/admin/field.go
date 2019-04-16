package admin

import (
	"github.com/cryo-management/api/models"
)

type FieldService struct{}

func (f *FieldService) Create(field *models.Field) error {
	err := field.Create()
	if err != nil {
		return err
	}

	translationService := new(TranslationService)
	err = translationService.Create(*field, field.ID)
	if err != nil {
		return err
	}

	return nil
}

func (f *FieldService) Load(field *models.Field, id string) error {
	err := field.Load(id)
	if err != nil {
		return err
	}

	return nil
}

func (f *FieldService) LoadAll(fields *models.Fields, schemaID string) error {
	err := fields.Load(schemaID)
	if err != nil {
		return err
	}

	return nil
}

func (f *FieldService) Delete(field *models.Field, id string) error {
	err := field.Delete(id)
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

	return nil
}

func (f *FieldService) DeleteBySchemaID(fields *models.Fields, id string) error {
	fieldService := new(FieldService)
	err := fieldService.LoadAll(fields, id)
	if err != nil {
		return err
	}

	for _, field := range *fields {
		err := fieldService.Delete(&field, field.ID)
		if err != nil {
			return err
		}
	}

	return nil
}
