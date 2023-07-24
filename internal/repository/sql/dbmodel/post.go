package dbmodel

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title      string      `json:"title"`
	Text       string      `json:"text"`
	Categories []*Category `json:"cats" gorm:"many2many:post_categories;"`
}
