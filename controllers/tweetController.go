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

func GetTweet(c *gin.Context)  {

    id := c.Param("id")
    var tweet models.Tweet
    initializers.DB.First(&tweet, "id = ?", id)
    //TODO validar que mandermos 404 cuando no encuentre un tweet

    c.HTML(http.StatusOK, "prueba.html", tweet)
}

