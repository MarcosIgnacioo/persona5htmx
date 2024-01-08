package controllers

import (
	"net/http"
	"github.com/MarcosIgnacioo/personahtmx/crud"
	"github.com/MarcosIgnacioo/personahtmx/helpers"
	"github.com/MarcosIgnacioo/personahtmx/initializers"
	"github.com/MarcosIgnacioo/personahtmx/models"
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

func RegisterUser(c *gin.Context)  {
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

    tweets := crud.GetAll() 
    session, _ := initializers.Store.Get(c.Request, "session")
    session.Values["user"] = user
    session.Save(c.Request, c.Writer)


    helpers.CreateSession(&user, c)
    data := models.Session { User: &user, Tweets: tweets, }
    c.HTML(http.StatusOK, "index.html", data)
}

func LogInUser(c *gin.Context)  {
    username := c.PostForm("username")
    password := c.PostForm("password")
    user := crud.Get(username)

    if helpers.CheckPasswordHash(password, user.Password) {

        tweets := crud.GetAll() //Tunear esto apra que soljjo de los primeros 50 
        helpers.CreateSession(user, c)
        data := models.Session { User: user, Tweets: tweets, }

        c.HTML(http.StatusOK, "index.html", data )
    } else {
        c.Status(400)
        return
    }
}

func LogOut(c *gin.Context) {
    session, _ := initializers.Store.Get(c.Request, "session")
    delete(session.Values, "user")
    session.Save(c.Request, c.Writer)
    return
}

func Count(c *gin.Context) {
}

func Test(c *gin.Context) {
    user, err := helpers.GetUserSession(c)
    if err != nil {
        c.HTML(http.StatusForbidden, "login.html", nil)
        return
    }
    c.HTML(http.StatusOK, "test.html", user)
}
