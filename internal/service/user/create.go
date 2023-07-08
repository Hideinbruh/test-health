package user

import (
	"fmt"
	"regexp"

	userModel "github.com/Hideinbruh/test-health/internal/model/user"
	"github.com/Hideinbruh/test-health/pkg/logger"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

func (s *service) Create(user *userModel.UserInfo) (int64, error) {
	if err := validateUserInfo(user); err != nil {
		logger.Infof("error: %s", err.Error())
		return 0, fmt.Errorf("validate: %w", err)
	}

	userId, err := s.repo.SaveUser(user)
	if err != nil {
		logger.Infof("error: %s", err.Error())
		return 0, err
	}

	logger.Infof("user %d created", userId)
	return userId, nil
}

func validateUserInfo(user *userModel.UserInfo) error {
	if err := validation.ValidateStruct(user,
		validation.Field(&user.Email, validation.Required, is.Email),
		validation.Field(&user.Password, validation.Required, validation.Length(6, 100)),
		validation.Field(&user.Username, validation.Required, validation.Match(regexp.MustCompile("^[A-Z]{2}$"))),
		validation.Field(&user.Phone, validation.Required, validation.Match(regexp.MustCompile("^[0-9]{11}$]")))); err != nil {
		return err
	}

	return nil
}
