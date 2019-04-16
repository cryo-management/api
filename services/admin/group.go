package admin

import (
	"github.com/cryo-management/api/models"
)

type GroupService struct{}

func (g *GroupService) Create(group *models.Group) error {
	err := group.Create()
	if err != nil {
		return err
	}

	translationService := new(TranslationService)
	err = translationService.Create(*group, group.ID)
	if err != nil {
		return err
	}

	return nil
}

func (g *GroupService) Load(group *models.Group, id string) error {
	err := group.Load(id)
	if err != nil {
		return err
	}

	return nil
}

func (g *GroupService) LoadAll(groups *models.Groups) error {
	err := groups.Load()
	if err != nil {
		return err
	}

	return nil
}

func (g *GroupService) Delete(group *models.Group, id string) error {
	err := group.Delete(id)
	if err != nil {
		return err
	}

	translationService := new(TranslationService)
	err = translationService.DeleteByStructureID(id)
	if err != nil {
		return err
	}

	groupUser := new(models.GroupUser)
	groupUser.GroupID = id

	groupService := new(GroupService)
	err = groupService.DeleteGoupUser(groupUser)
	if err != nil {
		return err
	}

	groupPermission := new(models.GroupPermission)
	groupPermission.GroupID = id

	err = groupService.DeletePermission(groupPermission)
	if err != nil {
		return err
	}

	return nil
}

func (g *GroupService) CreateGroupUser(groupUser *models.GroupUser) error {
	err := groupUser.Create()
	if err != nil {
		return err
	}

	return nil
}

func (g *GroupService) DeleteGoupUser(groupUser *models.GroupUser) error {
	err := groupUser.Delete()
	if err != nil {
		return err
	}

	return nil
}

func (g *GroupService) CreatePermission(groupPermission *models.GroupPermission) error {
	err := groupPermission.Create()
	if err != nil {
		return err
	}

	return nil
}

func (g *GroupService) DeletePermission(groupPermission *models.GroupPermission) error {
	err := groupPermission.Delete()
	if err != nil {
		return err
	}

	return nil
}
