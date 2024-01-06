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

func CreateTest( c *gin.Context ) {
    test := models.Prueba { Name: "prueba" }
    result := initializers.DB.Create(&test)
    if result.Error != nil {
        c.Status(400)
        return
    }
    c.HTML(http.StatusOK, "prueba.html", test)
}

