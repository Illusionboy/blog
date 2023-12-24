package service

import (
	"blog/common/global"
	"blog/models"
)

type ChannelService struct {
}

// 添加
func (u *ChannelService) AddChannel(channel models.Channel) int64 {
	return global.Db.Create(&channel).RowsAffected
}

// 删除
func (u *ChannelService) DelChannel(id int) int64 {
	return global.Db.Delete(&models.Channel{}, id).RowsAffected
}

// 修改
func (u *ChannelService) UpdateChannel(channel models.Channel) int64 {
	return global.Db.Updates(&channel).RowsAffected
}

// get
func (u *ChannelService) GetChannel(id int) models.Channel {
	var channel models.Channel
	global.Db.First(&channel, id)
	return channel
}

// get with slug
func (u *ChannelService) GetChannelBySlug(slug string) models.Channel {
	var channel models.Channel
	global.Db.Where("slug = ?", slug).First(&channel)
	return channel
}

// get Lists By Channel
func (u *ChannelService) GetChannelPostList(id uint64) []models.Post {
	postList := make([]models.Post, 0)

	global.Db.Where("channel_id = ?", id).Find(&postList)
	return postList
}

// get channel list
func (u *ChannelService) GetChannelList() []models.Channel {
	channelList := make([]models.Channel, 0)
	global.Db.Find(&channelList)
	return channelList
}
