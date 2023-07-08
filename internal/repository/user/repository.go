package user

import (
	userModel "github.com/Hideinbruh/test-health/internal/model/user"
)

type RepositoryUser interface {
	SaveUser(user *userModel.UserInfo) (int64, error)
}

type repositoryUser struct {
}

func NewRepositoryUser() *repositoryUser {
	return &repositoryUser{}
}

func (r *repositoryUser) SaveUser(user *userModel.UserInfo) (int64, error) {
	return 1, nil
}
