package backend

import (
	"changeme/backend/models"
)

type UserService struct{}

func (u *UserService) SaveUserInfo(user models.User) error {
	return nil;
}