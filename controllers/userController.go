package controllers

import (
	"net/http"

	"github.com/MarcosIgnacioo/personahtmx/crud"
	"github.com/MarcosIgnacioo/personahtmx/helpers"
	"github.com/MarcosIgnacioo/personahtmx/initializers"
	"github.com/MarcosIgnacioo/personahtmx/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func init()  {
    initializers.LoadEnvVariables()
    initializers.ConnectToDB()
}

func CreateTest( c *gin.Context ) {
    test := models.Prueba { Name: "prueba" }
    result := initializers.DB.Create(&test)
    if result.Error != nil {
        c.Status(400)
        return
    }
    c.HTML(http.StatusOK, "prueba.html", test)
}

func CreateUser(c *gin.Context)  {
    username := c.PostForm("username") 
    password := c.PostForm("password") 

    hashedPassword, err := helpers.HashPassword(password)
    if err != nil {
        c.Status(500)
        return
    }

    user := models.User { Username: username , Password: hashedPassword }

    newUser := initializers.DB.Create(&user)

    if newUser.Error != nil {
        c.Status(400)
        return
    }

    sessionStorage := sessions.Default(c)
    sessionStorage.Set("userId", user.ID)
    tweets := crud.GetAll() 

    helpers.CreateSession(&user, c)
    session := models.Session { User: &user, Tweets: tweets, }
    c.HTML(http.StatusOK, "index.html", session)
}

func LogInUser(c *gin.Context)  {
    username := c.PostForm("username")
    password := c.PostForm("password")
    user := crud.Get(username)

    if helpers.CheckPasswordHash(password, user.Password) {
        tweets := crud.GetAll()//Tunear esto apra que solo de los primeros 50 
        helpers.CreateSession(user, c)

        session := models.Session { User: user, Tweets: tweets, }

        c.HTML(http.StatusOK, "index.html", session )
    } else {
        c.Status(400)
        return
    }
}
