package controller

import (
	"blog/models"
	"blog/service"
	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var indexChannelService service.ChannelService
var indexPostService service.PostService

// 后台页面
func AdminIndex(c *gin.Context) {
	gintemplate.HTML(c, http.StatusOK, "index", nil)
}

// 前台页面
func Index(c *gin.Context) {
	channels := indexChannelService.GetChannelList()
	posts := indexPostService.GetPostList()
	page := c.Query("page")
	pageNumber, _ := strconv.Atoi(page)
	items, _ := GetItems(pageNumber, posts)
	posts = items
	gintemplate.HTML(c, http.StatusOK, "index", gin.H{"clist": channels, "posts": posts})
}

func GetItems(page int, posts []models.Post) ([]models.Post, error) {
	postsPerPage := 10
	tmpPosts := posts[page*postsPerPage : (page+1)*postsPerPage]
	var validPosts = make([]models.Post, 0)
	for _, post := range tmpPosts {
		if post.Title != "" {
			validPosts = append(validPosts, post)
		} else {
			break
		}
	}
	return validPosts, nil
}
