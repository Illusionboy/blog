package models

import (
	"blog/common/global"
	"fmt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       int    `gorm:"id;;primaryKey;AUTO_INCREMENT"`
	Username string `form:"username" gorm:"username;UNIQUE;NOT NULL"`
	Password []byte `form:"password" gorm:"password;NOT NULL"`
	Phone    string `form:"phone" gorm:"phone"`
	Email    string `form:"email" gorm:"email"`
	//Status   int    `gorm:"column:status;default:1;NOT NULL"`
	Status int `gorm:"column:status;NOT NULL"`
}

func (m *User) User() string {
	return "User"
}

func InitUser() {
	err := global.Db.AutoMigrate(&User{})
	if err != nil {
		fmt.Printf("mysql error: %s", err)
		return
	}
}
