package model

import (
	"github.com/ginblog/utils/errmsg"
	"gorm.io/gorm"
)

type Podcast struct {
	gorm.Model
	Title string `gorm:"type:varchar(100) not null" json:"title"`
	Url   string `gorm:"type:varchar(100)" json:"url"`
}

func CreatePodcast(data *Podcast) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 检查播客是否重复
func CheckPodcast(title string) int {
	var data *Podcast
	db.Where("title = ?", title).First(&data)
	if data != nil && data.Title == title {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除指定播客
func DeletePodcast(title string) int {
	var data *Podcast
	err := db.Where("title = ?", title).Delete(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
