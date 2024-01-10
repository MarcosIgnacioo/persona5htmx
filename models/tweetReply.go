package models

import "gorm.io/gorm"

type TweetReply struct {
    gorm.Model
    Author string `gorm:"column:author"`
    Content string `gorm:"column:content"`
    ParentTweetID uint `gorm:"column:parent_tweet_id"`
    ParentTweet Tweet `gorm:"foreignkey:ParentTweetID"`
}

func NewTweetReply(author string, content string, parentTweetID uint) TweetReply {
    return TweetReply{
        Author: author,
        Content: content,
        ParentTweetID: parentTweetID,
    }
}
