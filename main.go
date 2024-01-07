package main

import (
	"github.com/MarcosIgnacioo/personahtmx/controllers"
	"github.com/MarcosIgnacioo/personahtmx/initializers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/gorm"
	"github.com/gin-gonic/gin"
)

func main()  {

    r := gin.Default()
    r.LoadHTMLGlob("public/templates/*")
    r.Static("/assets", "./assets")
    store := gorm.NewStore(initializers.DB, true, []byte("secret"))
    r.Use(sessions.Sessions("authentication", store))

    // Routes
    r.GET("/", controllers.IndexView)
    r.GET("/register", controllers.RegisterView)
    r.GET("/login", controllers.LoginView)
    r.POST("/user", controllers.CreateUser)
    r.POST("/login-user", controllers.LogInUser)
    // Tweets
    r.POST("/tweet", controllers.CreateTweet)
    r.GET("/tweet/:id", controllers.GetTweet)
    r.Run() 
}
