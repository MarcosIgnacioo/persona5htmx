package main

import (
	"github.com/MarcosIgnacioo/personahtmx/controllers"
	"github.com/gin-gonic/gin"
)

func main()  {

    r := gin.Default()
    r.LoadHTMLGlob("public/templates/*")
    r.Static("/assets", "./assets")

    // Routes
    
    // Views
    r.GET("/", controllers.IndexView)
    r.GET("/register", controllers.RegisterView)
    r.GET("/login", controllers.LoginView)
    r.GET("/count", controllers.CountView)
    // User 
    r.POST("/user", controllers.RegisterUser)
    r.POST("/login-user", controllers.LogInUser)
    r.POST("/logout-user", controllers.LogOut)
    // Tweets
    r.POST("/tweet", controllers.CreateTweet)
    r.GET("/tweet/:id", controllers.GetTweet)
    r.POST("/incr", controllers.Count) 
    r.GET("/test", controllers.Test)
    r.Run() 
}

