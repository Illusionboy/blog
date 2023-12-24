package models

import (
	"blog/common/global"
	"fmt"
)

// Channel 数据库，用户数据映射模型
type Channel struct {
	ID    uint64 `form:"id" gorm:"primaryKey"`
	Title string `form:"title" gorm:"title"`
	// Slug用于构建 URL 的友好、可读性良好的字符串表示，路由用到的不包含空格的路径
	Slug    string `form:"slug" gorm:"slug"`
	Content string `form:"content" gorm:"content"`
	Status  int    `form:"status" gorm:"status"`
	// 用于对Channel进行排序
	Weight int `form:"weight" gorm:"weight"`
}

func (m *Channel) Channel() string {
	return "Channel"
}

func InitChannel() {
	err := global.Db.AutoMigrate(&Channel{})
	if err != nil {
		fmt.Printf("mysql error: %s", err)
		return
	}
}
