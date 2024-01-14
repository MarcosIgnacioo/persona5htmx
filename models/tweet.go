package models

import (
	"gorm.io/gorm"
)

type Tweet struct {
    gorm.Model
    Author        string   `gorm:"column:author"`
    Content       string   `gorm:"column:content"`
    Likes         uint     `gorm:"column:likes"`
    Retweets      uint     `gorm:"column:retweets"`
    LatestReplyID uint     
    LatestReply   string     `gorm:"default:NULL"`
}
