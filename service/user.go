package service

import (
	"blog/common/global"
	"blog/models"
)

type UserService struct {
}

// 添加
func (u *UserService) AddUser(user models.User) int64 {
	//var user_tmp models.User
	//result := global.Db.Where("username = ?", user.Username).First(&user_tmp)
	//if result.Error == nil {
	//	// 记录存在
	//	fmt.Println("Record exists:", user)
	//	return 0
	//} else if result.Error == gorm.ErrRecordNotFound {
	//	// 记录不存在
	//	fmt.Println("Record does not exist")
	//	return global.Db.Create(&user).RowsAffected
	//}
	return global.Db.Create(&user).RowsAffected
}

// 删除
func (u *UserService) DelUser(id int) int64 {
	return global.Db.Delete(&models.User{}, id).RowsAffected
}

// 修改
func (u *UserService) UpdateUser(user models.User) int64 {
	return global.Db.Where("username = ?", user.Username).Updates(&user).RowsAffected
}

// get
func (u *UserService) GetUser(id int) models.User {
	var user models.User
	global.Db.First(&user, id)
	return user
}

// get user list
func (u *UserService) GetUserList() []models.User {
	userList := make([]models.User, 0)
	global.Db.Find(&userList)
	return userList
}

// get admin list
func (u *UserService) GetAdminList() []models.User {
	userList := make([]models.User, 0)
	global.Db.Where("status = ?", 0).Find(&userList)
	return userList
}
