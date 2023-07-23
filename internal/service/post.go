package service

import (
	catlisthandler "blog-service-v3/internal/service/catListHandler"
)

type Post struct {
	Title      string
	Text       string
	Categories []string
}

func (a *postApp) CreatePost(title, text string, categories []string) error {
	return a.postRepo.CreatePost(title, text, categories)
}

func (a *postApp) AllPosts() ([]Post, error) {
	postRows, err := a.postRepo.AllPosts()
	if err != nil {
		return nil, err
	}

	posts := []Post{}
	for _, postRow := range postRows {
		posts = append(posts, Post{
			Title:      postRow.Title,
			Text:       postRow.Text,
			Categories: catlisthandler.CategoriesToList(postRow.Categories),
		})
	}

	return posts, nil
}

func (a *postApp) PagePosts(pageNumber, pageSize int) ([]Post, error) {
	postRows, err := a.postRepo.PagePosts(pageNumber, pageSize)
	if err != nil {
		return nil, err
	}

	posts := []Post{}
	for _, postRow := range postRows {
		posts = append(posts, Post{
			Title:      postRow.Title,
			Text:       postRow.Text,
			Categories: catlisthandler.CategoriesToList(postRow.Categories),
		})
	}

	return posts, nil
}

func (a *postApp) UpdatePost(postID int, title, text string, categories []string) error {
	return a.postRepo.UpdatePost(postID, title, text, categories)
}

func (a *postApp) DeletePost(postID int) error {
	return a.postRepo.DeletePost(postID)
}
