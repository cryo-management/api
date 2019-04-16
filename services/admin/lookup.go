package admin

import (
	"github.com/cryo-management/api/models"
)

type LookupService struct{}

func (l *LookupService) Create(lookup *models.Lookup) error {
	err := lookup.Create()
	if err != nil {
		return err
	}

	translationService := new(TranslationService)
	err = translationService.Create(*lookup, lookup.ID)
	if err != nil {
		return err
	}

	return nil
}

func (l *LookupService) Load(lookup *models.Lookup, id string) error {
	err := lookup.Load(id)
	if err != nil {
		return err
	}

	return nil
}

func (l *LookupService) LoadAll(lookups *models.Lookups) error {
	err := lookups.Load()
	if err != nil {
		return err
	}

	return nil
}

func (l *LookupService) Delete(lookup *models.Lookup, id string) error {
	err := lookup.Delete(id)
	if err != nil {
		return err
	}

	translationService := new(TranslationService)
	err = translationService.DeleteByStructureID(id)
	if err != nil {
		return err
	}

	return nil
}
