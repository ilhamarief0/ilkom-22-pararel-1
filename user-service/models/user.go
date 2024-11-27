package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	RoleID   int    `json:"role_id"`
}

type Role struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
