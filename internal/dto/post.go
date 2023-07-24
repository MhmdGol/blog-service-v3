package dto

type CreatePostRequest struct {
	Title string   `json:"title"`
	Text  string   `json:"text"`
	Cats  []string `json:"cats"`
}

type AllPostsResponse struct {
	Posts []Post `json:"posts"`
}

type Post struct {
	ID    uint     `json:"id"`
	Title string   `json:"title"`
	Text  string   `json:"text"`
	Cats  []string `json:"cats"`
}
