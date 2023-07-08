package user

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Age      int    `json:"age"`
}

type UserInfo struct {
	Username string
	Password string
	Email    string
	Phone    string
}
