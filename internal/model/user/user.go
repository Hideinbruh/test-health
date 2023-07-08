package user

type HandlerUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type User struct {
	Username string
	Password string
	Email    string
	Phone    string
}
