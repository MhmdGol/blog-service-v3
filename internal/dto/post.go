package dto

type CreatePostRequest struct {
	Title string   `json:"title" validate:"required,min=5,max=100"`
	Text  string   `json:"text" validate:"required,min=10,max=500"`
	Cats  []string `json:"cats" validate:"required,dive,required,unique"`
}

type AllPostsResponse struct {
	Posts []Post `json:"posts"`
}

type Post struct {
	ID    string   `json:"id"`
	Title string   `json:"title"`
	Text  string   `json:"text"`
	Cats  []string `json:"cats"`
}
