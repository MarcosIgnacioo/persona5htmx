package controllers

import (
	"net/http"
	"github.com/MarcosIgnacioo/personahtmx/helpers"
	"github.com/MarcosIgnacioo/personahtmx/models"
	"github.com/gin-gonic/gin"
)

func RegisterView(c *gin.Context)  {
    c.HTML(http.StatusOK, "register.html", nil)
}

func IndexView(c *gin.Context)  {
    // Solucion a todo esto, dejar todos los html como antes, pero en los handlers checar por la session y si no hay una sesion redirigir al login pero pues si hay una sesion seguir con todos los procesos
    user, err := helpers.GetUserSession(c)
    if err != nil {
        c.HTML(http.StatusOK, "login.html", nil)
    }
    tweetsHome := helpers.LoadTweets(c)
    session := models.Session { User: user.Username, Render: &tweetsHome, }

    c.HTML(http.StatusOK, tweetsHome.Template , session)
}


func LoginView(c *gin.Context)  {
    c.HTML(http.StatusOK, "login.html", nil)
}
func CountView(c *gin.Context)  {
    c.HTML(http.StatusOK, "test.html", nil)
}

func CheckError(err error) {
    if err != nil {
        panic("Hubo un error")
    }
}
