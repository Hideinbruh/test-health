package converter

import "awesomeProject3/internal/model/user"

func ToCreateService(userHandler user.HandlerUser) *user.User {
	return &user.User{
		Username: userHandler.Username,
		Password: userHandler.Password,
		Email:    userHandler.Email,
		Phone:    userHandler.Phone,
	}

}
