package controllers

import (
	"net/http"
	"strconv"

	"github.com/MarcosIgnacioo/personahtmx/helpers"
	"github.com/MarcosIgnacioo/personahtmx/initializers"
	"github.com/MarcosIgnacioo/personahtmx/models"
	"github.com/gin-gonic/gin"
)

func init()  {
    initializers.LoadEnvVariables()
    initializers.ConnectToDB()
}

func CreateTweet(c *gin.Context)  {
    tweetContent := c.PostForm("tweet")
    author := c.PostForm("username")

    tweet := models.Tweet { 
        Author: author,
        Content: tweetContent,
    }

    result := initializers.DB.Create(&tweet)

    if result.Error != nil {
        c.Status(400)
        return
    }
    c.HTML(http.StatusOK, "tweet.html", tweet)
}

func ReplyTweet(c *gin.Context)  {
    id := c.Param("id")
    idR, err := strconv.Atoi(id)
    replyContent := c.PostForm("reply-content")
    replyAuthor, _ := helpers.GetUserSession(c)

    if err != nil {
        c.String(http.StatusBadGateway, "Pop") 
    }

    reply := models.NewTweetReply(replyAuthor.Username, replyContent, uint(idR))

    var tweet models.Tweet
    initializers.DB.First(&tweet, "id = ?", id)


    initializers.DB.Create(&reply)

    tweet.LatestReplyID = reply.ID

    initializers.DB.Save(tweet)
    c.String(http.StatusOK, reply.Content) 
}


type TweetPlus struct {
    Author        string `gorm:"column:author_parent"`
    Content       string `gorm:"column:content_parent"`
    AuthorReply   string `gorm:"column:author_reply"`
    ContentReply  string `gorm:"column:content_reply"`
}
type TweetResponse struct {
    ContentReply   string `json:"content_reply"`
    AuthorReply    string `json:"author_reply"`
}

type TweetResult struct {
	ContentParent  string   `json:"content_parent"`
	AuthorParent   string   `json:"author_parent"`
	Replies        []string `json:"replies"`
}

func GetTweet(c *gin.Context)  {

    id := c.Param("id")
    var tweetResult TweetResult
    initializers.DB.Table("tweets").
        Select("tweets.content as content_parent, tweets.author as author_parent, GROUP_CONCAT(CONCAT(tweet_replies.content, ' - ', tweet_replies.author)) AS replies").
        Joins("RIGHT JOIN tweet_replies ON tweet_replies.parent_tweet_id = tweets.id").
        Where("tweets.id = ?", id).
        Group("tweets.content, tweets.author").
        Order("tweets.updated_at DESC").
        Scan(&tweetResult)

    //TODO validar que mandermos 404 cuando no encuentre un tweet

    c.HTML(http.StatusOK, "test.html", tweetResult)
}

