package controller

//func LoginOK(c *gin.Context) {
//	username := c.PostForm("username")
//	password := utils.EncryptHash(c.PostForm("password"))
//	fmt.Println(username)
//	u := dao.Mgr.Login(username)
//
//	if u.Username == "" {
//		c.HTML(http.StatusOK, "login.html", "用户不存在！")
//	} else {
//		if u.Password != password {
//			fmt.Println("密码错误！")
//			c.HTML(http.StatusOK, "login.html", "密码错误！")
//		} else {
//			fmt.Println("登录成功")
//			c.Redirect(http.StatusMovedPermanently, "/")
//		}
//	}
//	c.Redirect(http.StatusMovedPermanently, "/")
//}
