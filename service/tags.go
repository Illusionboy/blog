package service

import (
	"blog/common/global"
	"blog/models"
)

type TagsService struct {
}

// 添加
func (p *TagsService) AddTags(tags models.Tags) int64 {
	return global.Db.Table("tags").Create(&tags).RowsAffected
}

// 删除
func (p *TagsService) DeleteTags(id int) int64 {
	return global.Db.Delete(&models.Tags{}, id).RowsAffected
}

// 修改
func (p *TagsService) UpdateTags(tags models.Tags) int64 {
	return global.Db.Updates(&tags).RowsAffected
}

// get
func (p *TagsService) GetTags(id int) models.Tags {
	var tags models.Tags
	global.Db.First(tags, id)
	return tags
}

// get tags list
func (p *TagsService) GetTagsList() []models.Tags {
	tagsList := make([]models.Tags, 0)
	global.Db.Find(&tagsList)
	return tagsList
}
