package user

import (
	userModel "awesomeProject3/internal/model/user"
)

type RepositoryUser interface {
	SaveUser(user *userModel.User) error
}

type repositoryUser struct {
}

func NewRepositoryUser() *repositoryUser {
	return &repositoryUser{}
}

func (r *repositoryUser) SaveUser(user *userModel.User) error {
	return nil
}
