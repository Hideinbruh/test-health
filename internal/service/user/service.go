package user

import (
	"github.com/Hideinbruh/test-health/internal/model/user"
	userRepository "github.com/Hideinbruh/test-health/internal/repository/user"
)

var _ Service = (*service)(nil)

type Service interface {
	Create(user *user.UserInfo) (int64, error)
}

type service struct {
	repo userRepository.RepositoryUser
}

func NewService() *service {
	return &service{}
}
