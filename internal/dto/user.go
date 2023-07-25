package dto

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"user"`
	Password string `json:"pass"`
}
