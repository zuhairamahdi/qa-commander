package models

type User struct {
	ID           uint   `json:"id"`
	Username     string `json:"username"`
	PasswordHash string `json:"password_hash"`
	Salt         string `json:"salt"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	Active       bool   `json:"active"`
	Email        string `json:"email"`
	Role         string `json:"role"`
}
