package models

import (
	"blog/common/global"
	"fmt"
	"gorm.io/gorm"
)

type Tags struct {
	gorm.Model
	ID   int    `gorm:"id;primaryKey;AUTO_INCREMENT"`
	Name string `gorm:"name;NOT NULL"`
}

func (m *Tags) Tags() string {
	return "Tags"
}

func InitTags() {
	err := global.Db.AutoMigrate(&Tags{})
	if err != nil {
		fmt.Printf("mysql error: %s", err)
		return
	}
}
