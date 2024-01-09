package controllers

import (
	"net/http"
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
    c.HTML(http.StatusOK, "tweet-form.html", tweet)
    c.HTML(http.StatusOK, "oob-tweet.html", tweet)
}

func GetTweet(c *gin.Context)  {

    id := c.Param("id")
    var tweet models.Tweet
    initializers.DB.First(&tweet, "id = ?", id)
    //TODO validar que mandermos 404 cuando no encuentre un tweet

    c.HTML(http.StatusOK, "prueba.html", tweet)
}
