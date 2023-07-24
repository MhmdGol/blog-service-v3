package model

type Post struct {
	ID         ID
	Title      string
	Text       string
	Categories []string
}
