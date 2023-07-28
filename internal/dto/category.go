package dto

type CreateCategoryRequest struct {
	Name string `json:"name" validate:"required,min=3,max=50,alphanum"`
}

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
