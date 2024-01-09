package models

import "gorm.io/gorm"

type Tweet struct {
    gorm.Model
    ParentTweetID *uint    `gorm:"column:parent_tweet_id;default:null"`
    Replies       []Tweet  `gorm:"foreignKey:ParentTweetID"`
    Author        string   `gorm:"column:author"`
    Content       string   `gorm:"column:content"`
    Likes         uint     `gorm:"column:likes"`
    Retweets      uint     `gorm:"column:retweets"`
}
