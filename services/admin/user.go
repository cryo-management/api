package admin

import (
	"github.com/cryo-management/api/models"
)

type UserService struct{}

func (u *UserService) Create(user *models.User) error {
	err := user.Create()
	if err != nil {
		return err
	}

	return nil
}

func (u *UserService) Load(user *models.User, id string) error {
	err := user.Load(id)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserService) LoadAll(users *models.Users) error {
	err := users.Load()
	if err != nil {
		return err
	}

	return nil
}

func (u *UserService) Delete(user *models.User, id string) error {
	err := user.Delete(id)
	if err != nil {
		return err
	}

	groupUser := new(models.GroupUser)
	groupUser.UserID = id

	groupService := new(GroupService)
	err = groupService.DeleteGoupUser(groupUser)
	if err != nil {
		return err
	}

	return nil
}
