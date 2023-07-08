package user

import (
	userModel "awesomeProject3/internal/model/user"
	"awesomeProject3/internal/repository/user"
)

type Service interface {
	Create(user *userModel.User) (int64, error)
}

type service struct {
	Repo user.RepositoryUser
}

func NewService() *service {
	return &service{}
}
