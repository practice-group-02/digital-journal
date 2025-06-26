package models

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	PasswordHashed  string `json:"password"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
}

func NewUser() *User {
	return &User{}
}
