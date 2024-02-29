package models

const (
	RoleClient = "client"
	RoleAdmin  = "admin"
)

type User struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Password []byte `json:"-"`
	Role     string `json:"role"`
}
