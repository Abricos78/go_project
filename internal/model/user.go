package user

type User struct {
	Username string `json:"username"`
	Age      uint8  `json:"age"`
	id       string
	Email    string `json:"email"`
	Password string `json:"password"`
	isAdmin  bool
}

func (u *User) GetIsAdmin() bool {
	return u.isAdmin
}

func (u *User) SetIsAdmin(isAdmin bool) {
	u.isAdmin = isAdmin
}
