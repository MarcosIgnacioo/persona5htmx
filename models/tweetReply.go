package models

import "gorm.io/gorm"

type TweetReply struct {
    gorm.Model
    Author string `gorm:"column:author"`
    Content string `gorm:"column:content"`
    ParentId uint `gorm:"column:parent_id"` 
    ParentTweet Tweet `gorm:"column:parent_tweet"` 
}
