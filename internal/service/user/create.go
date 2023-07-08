package user

import (
	userModel "awesomeProject3/internal/model/user"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"regexp"
)

func (s *service) Create(user *userModel.User) (int64, error) {
	Validate(user)
	var userId int64 = 1
	if err := s.Repo.SaveUser(user); err != nil {
		return 0, err
	}
	return userId, nil
}

func Validate(user *userModel.User) error {
	err := validation.ValidateStruct(
		user, validation.Field(&user.Email, validation.Required, is.Email),
		validation.Field(&user.Password, validation.Required, validation.Length(6, 100)),
		validation.Field(&user.Username, validation.Required, validation.Match(regexp.MustCompile("^[A-Z]{2}$"))),
		validation.Field(&user.Phone, validation.Required, validation.Match(regexp.MustCompile("^[0-9]{11}$]"))))
	if err != nil {
		return err
	}
	return nil
}
