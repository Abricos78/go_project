package user

type User struct {
	Username string `json:"username"`
	Age      uint8  `json:"age"`
	id       string
	Email    string `json:"email"`
	isAdmin  bool
}

func (u *User) GetIsAdmin() bool {
	return u.isAdmin
}

func (u *User) SetIsAdmin(isAdmin bool) {
	u.isAdmin = isAdmin
}
