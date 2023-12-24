package controller

import (
	"blog/common/global"
	"blog/models"
	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SearchKeywords(c *gin.Context) {
	channels := indexChannelService.GetChannelList()
	// 从查询参数中获取搜索关键词
	keyword := c.Query("keywords")
	var posts []models.Post
	global.Db.Where("title LIKE ? OR tags LIKE ?", "%"+keyword+"%", "%"+keyword+"%").Find(&posts)
	gintemplate.HTML(c, http.StatusOK, "index", gin.H{"clist": channels, "posts": posts})
}
