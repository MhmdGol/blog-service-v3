package dto

type AuthToken struct {
	Token string `json:"token" validate:"required,regexp=^[a-zA-Z0-9]+$"`
}
