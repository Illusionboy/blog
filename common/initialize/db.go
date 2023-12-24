package initialize

import (
	"blog/common/global"
	"blog/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

// Mysql 配置MySQL数据库
func Mysql() {
	m := global.Config.Mysql
	var dsn = fmt.Sprintf("%s:%s@%s", m.Username, m.Password, m.Url)
	//fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		fmt.Printf("mysql error: %s", err)
		return
	}

	sqlDb, err := db.DB()
	if err != nil {
		fmt.Printf("mysql error: %s", err)
	}
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxIdleTime(time.Hour)
	global.Db = db

	models.InitChannel()
	models.InitComment()
	models.InitPost()
	models.InitTags()
	models.InitUser()
}
