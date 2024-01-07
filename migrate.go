package main

import (
	"github.com/MarcosIgnacioo/personahtmx/initializers"
	"github.com/MarcosIgnacioo/personahtmx/models"
)

func init()  {
    // Nos aseguramos de estar conectados a la DB no se que pasa si quitamos esto pero pues luego lo averiguamos
    initializers.LoadEnvVariables()
    initializers.ConnectToDB()
    initializers.DB.AutoMigrate(&models.Prueba{}, models.User{}, models.Tweet{})
    initializers.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.Prueba{})
}
