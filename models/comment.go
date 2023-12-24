package models

import (
	"blog/common/global"
	"fmt"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	PostID         int    `gorm:"column:post_id;NOT NULL"`
	CommentContent string `gorm:"column:post_id;NOT NULL"`
}

func (m *Comment) Comment() string {
	return "Comment"
}

func InitComment() {
	err := global.Db.AutoMigrate()
	if err != nil {
		fmt.Printf("mysql error: %s", err)
		return
	}
}
