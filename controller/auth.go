package controller

import (
	"blog/common/utils"
	"blog/models"
	"blog/service"
	"bytes"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Accounts map[string][]byte

var (
	adminAccounts = Accounts{}
	userAccounts  = Accounts{}
	RevokedTokens = make(map[string]bool) // 使用撤销列表实现 Token 失效
)

// <==================================================登录处理相关函数=================================================>
func GoLogin(c *gin.Context) {
	gintemplate.HTML(c, http.StatusOK, "login", nil)
}

//func AuthLogin() gin.HandlerFunc {
//	userList := make([]models.User, 2)
//	userService := service.UserService{}
//	userList = userService.GetUserList()
//	fmt.Println("In userService.GetUserList():")
//	//fmt.Printf("Your Input is            %x\n", password)
//	fmt.Printf("Root's Password is %x\n", userList[1])
//
//	userAccounts := Accounts{}
//	// 遍历结构体切片并构建 map
//	for _, user := range userList {
//		// 使用结构体中的 Name 作为键，Email 作为值
//		userAccounts[user.Username] = user.Password
//	}
//	return authLoginHandler(userAccounts)
//}

func AuthLogin(c *gin.Context) {
	userList := make([]models.User, 2)
	userService := service.UserService{}
	userList = userService.GetUserList()

	for _, user := range userList {
		// 使用结构体中的 Name 作为键，Email 作为值
		// 使用结构体中的 Name 作为键，Email 作为值
		userAccounts[user.Username] = user.Password
	}

	username := c.PostForm("username")
	password := utils.EncryptPassword(c.PostForm("password"))
	fmt.Println(username)
	if !authenticateUser(username, password, userAccounts) {
		utils.RespondWithError(401, "Unauthorized", c)
		return
	}
	// 生成token
	if tokenString, err := GenToken(username); err == nil {
		//c.JSON(http.StatusOK, gin.H{"jwtToken": tokenString})
		// 设置Cookie的过期时间为2小时
		expiration := time.Now().Add(time.Hour * 2)
		maxAge := int(expiration.Sub(time.Now()).Seconds())

		// 将jwtToken写入Cookie
		cookie := &http.Cookie{
			Name:    "jwtToken",
			Value:   tokenString,
			Expires: expiration,
			MaxAge:  maxAge,
			Secure:  false,
		}
		http.SetCookie(c.Writer, cookie)
		c.Redirect(http.StatusMovedPermanently, "/")
		return
	}
	utils.RespondWithError(401, "BadAuth", c)
}

func AuthAdmin() gin.HandlerFunc {
	adminList := make([]models.User, 2)
	userService := service.UserService{}
	adminList = userService.GetAdminList()

	//userAccounts := gin.Accounts{}
	// 遍历结构体切片并构建 map
	for _, admin := range adminList {
		// 使用结构体中的 Name 作为键，Email 作为值
		adminAccounts[admin.Username] = admin.Password
	}
	return authLoginHandler(adminAccounts)
}

func authLoginHandler(accounts Accounts) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.PostForm("username")
		password := utils.EncryptPassword(c.PostForm("password"))
		fmt.Println(username)

		if !authenticateUser(username, password, accounts) {
			utils.RespondWithError(401, "Unauthorized", c)
			return
		}
		// 生成token
		if tokenString, err := GenToken(username); err == nil {
			//c.JSON(http.StatusOK, gin.H{"jwtToken": tokenString})
			// 设置Cookie的过期时间为2小时
			expiration := time.Now().Add(time.Hour * 2)
			maxAge := int(expiration.Sub(time.Now()).Seconds())

			// 将jwtToken写入Cookie
			cookie := &http.Cookie{
				Name:    "jwtToken",
				Value:   tokenString,
				Expires: expiration,
				MaxAge:  maxAge,
				Secure:  false,
			}
			http.SetCookie(c.Writer, cookie)
			c.Redirect(http.StatusMovedPermanently, "/")
			return
		}
		utils.RespondWithError(401, "BadAuth", c)
	}
}

func authenticateUser(username string, password []byte, accounts Accounts) bool {
	// fetch user from accounts.
	for usr, pswd := range accounts {
		//fmt.Printf("Check %s\n", usr)
		if usr == username {
			if bytes.Equal(password, pswd) {
				return true
			} else {
				fmt.Println("Wrong Password")
				//fmt.Printf("Your Input is            %x\n", password)
				//fmt.Printf("YourCorrect Password is %x\n", pswd)
				return false
			}
		}
	}
	//fmt.Println("Wrong Username")
	return false
}

func ChangePasswd(c *gin.Context) {
	userList := make([]models.User, 0)
	userService := service.UserService{}
	userList = userService.GetUserList()

	// 遍历结构体切片并构建 map
	for _, user := range userList {
		// 使用结构体中的 Name 作为键，Email 作为值
		userAccounts[user.Username] = user.Password
	}

	username := c.PostForm("username")
	password := utils.EncryptPassword(c.PostForm("password"))
	comfirmPassword := utils.EncryptPassword(c.PostForm("confirmPassword"))
	if !bytes.Equal(password, comfirmPassword) {
		c.JSON(http.StatusBadGateway, gin.H{"error": "Please input same password!"})
		return
	}
	// fetch user from accounts.
	for usr, _ := range userAccounts {
		//fmt.Printf("Check %s\n", usr)
		if usr == username {
			//fmt.Printf("Found %s\n", usr)
			user := models.User{
				Username: usr,
				Password: password,
			}
			//fmt.Printf("Your Input is            %x\n", user.Password)
			//fmt.Printf("Your old Password is %x\n", oldPasswd)
			userService.UpdateUser(user)
			tokenString, err := c.Cookie("jwtToken")
			if err != nil {
				utils.RespondWithError(401, "jwtToken err", c)
			}
			// 将旧 Token 添加到撤销列表中
			RevokedTokens[tokenString] = true
			c.Redirect(http.StatusMovedPermanently, "/login")
			return
		}
	}
	utils.RespondWithError(401, "Username Doesn't Exists", c)
}

func LogOut(c *gin.Context) {
	tokenString, err := c.Cookie("jwtToken")
	if err != nil {
		utils.RespondWithError(401, "jwtToken err", c)
	}
	// 将旧 Token 添加到撤销列表中
	RevokedTokens[tokenString] = true
	c.Redirect(http.StatusMovedPermanently, "/login")
	return
}

// <==================================================JWT处理相关函数=================================================>
type JWTClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

const jwtDuration = time.Hour * 2

var secretKey = []byte("Xlog-Token")

func GenToken(username string) (string, error) {
	// Create the Claims
	claims := JWTClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(jwtDuration).Unix(),
			Issuer:    "Xlog",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func ParseToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		fmt.Printf("%v %v", claims.Username, claims.StandardClaims.ExpiresAt)
		return claims, nil
	}
	return nil, err
}
