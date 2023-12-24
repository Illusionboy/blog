package controller

import (
	"blog/models"
	"fmt"
	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// channel list
func ListChannel(c *gin.Context) {
	c2 := ChannelServiceInstance.GetChannelList()
	gintemplate.HTML(c, http.StatusOK, "channel/list", gin.H{"clist": c2})
}

// view channel ADMIN
func ViewChannel(c *gin.Context) {
	sid, r := c.GetQuery("id")
	var chann models.Channel
	if r {
		// 如果r为true，则获取id值；如果id为空，为创建频道
		id, _ := strconv.Atoi(sid)
		chann = ChannelServiceInstance.GetChannel(id)
	}
	// **补充channel/view.html，频道详情
	gintemplate.HTML(c, http.StatusOK, "channel/view", gin.H{"channel": chann})
}

// View Channel With slug USER
func ChannelView(c *gin.Context) {
	slug := c.Param("slug")
	channel := ChannelServiceInstance.GetChannelBySlug(slug)

	posts := ChannelServiceInstance.GetChannelPostList(channel.ID)
	c2 := ChannelServiceInstance.GetChannelList()

	gintemplate.HTML(c, http.StatusOK, "/channel/list", gin.H{"channel": channel, "posts": posts, "clist": c2})
}

// delete channel ADMIN
func DeleteChannel(c *gin.Context) {
	sid, _ := c.GetQuery("id")
	id, _ := strconv.Atoi(sid)
	ChannelServiceInstance.DelChannel(id)
	c.Redirect(http.StatusFound, "/admin/channel/list")
}

// 添加或更新 ADMIN
func SaveChannel(c *gin.Context) {
	var chann models.Channel
	if err := c.ShouldBind(&chann); err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	chann.Status, _ = strconv.Atoi(c.PostForm("status"))

	id, _ := c.GetPostForm("id")
	if id != "0" {
		ChannelServiceInstance.UpdateChannel(chann)
	} else {
		ChannelServiceInstance.AddChannel(chann)
	}
	c.Redirect(http.StatusFound, "/admin/channel/list")
}
