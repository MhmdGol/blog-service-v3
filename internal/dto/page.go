package dto

type PageSize struct {
	Size int `json:"size" validate:"required,gte=1,lte=100"`
}
