package initialize

import (
	"blog/common/global"
	"blog/common/middleware"
	"blog/controller"
	"fmt"
	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router() {
	engine := gin.Default()

	// 静态资源请求映射
	engine.Static("/assets", "./assets")
	engine.StaticFS("/static", http.Dir("static"))
	engine.StaticFS("/views/images", http.Dir("static/images"))

	engine.HTMLRender = gintemplate.New(gintemplate.TemplateConfig{
		Root:         "templates/frontend",
		Extension:    ".html",
		Master:       "layouts/master",
		DisableCache: true,
	})

	// 前台
	// 上方展示channel，下方展示Post
	engine.GET("/", controller.Index)
	engine.GET("/login", controller.GoLogin)
	engine.POST("/login", controller.AuthLogin)
	engine.GET("/post/:id", controller.PostDetail)
	// 添加channel路由
	engine.GET("/channel/:slug", controller.ChannelView)
	engine.GET("/s", controller.SearchKeywords)

	mw := gintemplate.NewMiddleware(gintemplate.TemplateConfig{
		Root:         "templates/backend",
		Extension:    ".html",
		Master:       "layouts/master",
		DisableCache: true,
	})

	// 后台管理员前端接口
	web := engine.Group("/admin", middleware.AuthMiddleware(), mw)

	{
		// index
		web.GET("/", controller.AdminIndex)
		// 修改密码
		web.GET("chpasswd", controller.GoLogin)
		web.POST("chpasswd", controller.ChangePasswd)
		// 退出登录
		web.GET("logout", controller.LogOut)
	}

	{
		// channel
		// 用户登录API
		// **添加修改和删除按钮【html】
		web.GET("channel/list", controller.ListChannel)
		// **频道详情/修改【html】
		web.GET("channel/view", controller.ViewChannel)
		web.POST("channel/save", controller.SaveChannel)
		web.GET("channel/del", controller.DeleteChannel)
	}

	{
		// channel
		// 用户登录API
		web.GET("post/list", controller.ListPost)
		// **推文详情/修改，添加markdownEdit【html】
		web.GET("post/view", controller.ViewPost)
		web.POST("post/save", controller.SavePost)
		web.GET("post/del", controller.DeletePost)

		web.POST("post/upload", controller.UploadThumbnails)
	}

	// 启动，监听端口
	post := fmt.Sprintf(":%s", global.Config.Server.Post)
	if err := engine.Run(post); err != nil {
		fmt.Printf("Server start error: %s", err)
	}
}
