package converter

import "github.com/Hideinbruh/test-health/internal/model/user"

func ToCreateService(userHandler user.UserRequest) *user.UserInfo {
	return &user.UserInfo{
		Username: userHandler.Username,
		Password: userHandler.Password,
		Email:    userHandler.Email,
		Phone:    userHandler.Phone,
	}

}
