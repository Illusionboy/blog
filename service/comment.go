package service

import (
	"blog/common/global"
	"blog/models"
)

type CommentService struct {
}

// 添加
func (u *CommentService) AddComment(comment models.Comment) int64 {
	return global.Db.Create(&comment).RowsAffected
}

// 删除
func (u *CommentService) DelComment(id int) int64 {
	return global.Db.Delete(&models.Comment{}, id).RowsAffected
}

// 修改
func (u *CommentService) UpdateComment(comment models.Comment) int64 {
	return global.Db.Updates(&comment).RowsAffected
}

// get
func (u *CommentService) GetComment(id int) models.Comment {
	var comment models.Comment
	global.Db.First(&comment, id)
	return comment
}

// get comment list
func (u *CommentService) GetCommentList() []models.Comment {
	commentList := make([]models.Comment, 0)
	global.Db.Find(&commentList)
	return commentList
}
