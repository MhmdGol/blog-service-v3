package sqlmodel

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name  string
	Posts []*Post `gorm:"many2many:post_categories;"`
}
