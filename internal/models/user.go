package models

type User struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	PasswordHash string `json:"password_hash"`
	Role         string `json:"role"`
}

type UserCreateForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
