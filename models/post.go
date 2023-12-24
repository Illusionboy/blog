package models

import (
	"blog/common/global"
	"fmt"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	ID        int    `gorm:"id;;primaryKey"`
	Title     string `form:"title" gorm:"title;NOT NULL"`
	Thumbnail string `form:"thumbnail" gorm:"thumbnail"`
	Summary   string `form:"summary" gorm:"summary"`
	Content   string `form:"content" gorm:"content;NOT NULL"`
	Author    string `form:"author" gorm:"author"`
	ChannelId int    `form:"channelId" gorm:"channel_id"`
	//Comments  int    `form:"comments" gorm:"comments"`
	Favors int `form:"favors" gorm:"favors"`
	//Featured  int    `form:"featured" gorm:"featured"`
	Status int `form:"status" gorm:"status"`
	// **tags应为ID类型
	Tags  string `form:"tags" gorm:"tags"`
	Views int    `form:"views" gorm:"views"`
	//Weight    string `form:"weight" gorm:"weight"`
	//Url       string `form:"url" gorm:"url"`
}

func (m *Post) Post() string {
	return "Post"
}

func InitPost() {
	err := global.Db.AutoMigrate(&Post{})
	if err != nil {
		fmt.Printf("mysql error: %s", err)
		return
	}
}
