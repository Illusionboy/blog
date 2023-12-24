package controller

import (
	"blog/common/stripmd"
	"blog/models"
	"fmt"
	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
	"github.com/google/uuid"
	"github.com/spf13/cast"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

// post list
func ListPost(c *gin.Context) {
	plist := PostServiceInstance.GetPostList()
	gintemplate.HTML(c, http.StatusOK, "post/list", gin.H{"plist": plist})
}

// view post
func ViewPost(c *gin.Context) {
	sid, r := c.GetQuery("id")
	var post models.Post
	if r {
		// 如果有值，转换；否则创建
		id, _ := strconv.Atoi(sid)
		post = PostServiceInstance.GetPost(id)
	} else {
		post.Author = c.MustGet(gin.AuthUserKey).(string)
	}
	clist := ChannelServiceInstance.GetChannelList()
	// author_id**
	gintemplate.HTML(c, http.StatusOK, "post/view", gin.H{
		"post":  post,
		"clist": clist,
	})
}

func DeletePost(c *gin.Context) {
	sid, _ := c.GetQuery("id")
	id, _ := strconv.Atoi(sid)
	PostServiceInstance.DeletePost(id)
	c.Redirect(http.StatusFound, "/admin/post/list")
}

func PostDetail(c *gin.Context) {
	sid := c.Param("id")
	id := cast.ToInt(sid)
	post := PostServiceInstance.GetPost(id)

	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	parsered := parser.NewWithExtensions(extensions)

	md := []byte(post.Content)
	md = markdown.NormalizeNewlines(md)
	html := markdown.ToHTML(md, parsered, nil)

	c2 := ChannelServiceInstance.GetChannelList()

	gintemplate.HTML(c, http.StatusOK, "/post/detail", gin.H{"clist": c2, "post": post, "content": template.HTML(html)})

}

// 上传封面
func UploadThumbnails(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "没有文件",
		})
		return
	}

	ext := filepath.Ext(file.Filename)
	newFileName := uuid.New().String() + ext
	// The file is received, so let's save it
	pwd, _ := os.Getwd()
	filePath := fmt.Sprintf("%s/static/thumbnails/%s", pwd, newFileName)
	relativeFilePath := fmt.Sprintf("/static/thumbnails/%s", newFileName)

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "不能保存文件",
		})
		fmt.Printf("err: v\n", err)
		return
	}
	c.String(http.StatusOK, relativeFilePath)
}

// 添加或者更新
func SavePost(c *gin.Context) {
	var post models.Post

	//sid := c.PostForm("id")
	title := c.PostForm("title")
	thumbnail := c.PostForm("thumbnail")
	tags := c.PostForm("tags")
	content := c.PostForm("content")
	channelId := c.PostForm("channelId")
	author := c.PostForm("author")

	// 摘要， 去掉markdown格式
	summary := stripmd.Strip(content)
	l := len(summary)
	if l >= 200 {
		summary = summary[0:200]
	} else {
		summary = summary[0:l]
	}

	//id, _ := strconv.Atoi(sid)
	id := cast.ToInt(c.PostForm("id"))
	post.ID = id
	post.Title = title
	post.Thumbnail = thumbnail
	post.Tags = tags
	post.Content = content
	post.Author = author // 实现注册登录后，动态获得
	post.ChannelId = cast.ToInt(channelId)
	fmt.Printf("ChannelID is %v\n", post.ChannelId)

	if err := c.ShouldBind(&post); err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	if id != 0 {
		PostServiceInstance.UpdatePost(post)
	} else {
		PostServiceInstance.AddPost(post)
	}
	c.Redirect(http.StatusFound, "/admin/post/list")
}
