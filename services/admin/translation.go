package admin

import (
	"github.com/cryo-management/api/models"
)

type TranslationService struct{}

func (t *TranslationService) Create(obj interface{}, id string) error {
	translation := new(models.Translation)
	err := translation.Create(id, obj)
	if err != nil {
		return err
	}

	return nil
}

func (t *TranslationService) DeleteByStructureID(id string) error {
	translation := new(models.Translation)
	err := translation.DeleteByStructureID(id)
	if err != nil {
		return err
	}

	return nil
}
