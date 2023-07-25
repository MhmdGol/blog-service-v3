package dto

type CreateCategoryRequest struct {
	Name string `json:"name"`
}

type Category struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
