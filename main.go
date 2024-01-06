package main

import (
	"github.com/MarcosIgnacioo/personahtmx/controllers"
	"github.com/gin-gonic/gin"
)

func main()  {
    r := gin.Default()
    r.LoadHTMLGlob("public/templates/*")
    r.Static("/assets", "./assets")

    r.GET("/", controllers.RegisterView)
    r.Run() 
}
