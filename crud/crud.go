package crud

import (
	"github.com/MarcosIgnacioo/personahtmx/initializers"
	"github.com/MarcosIgnacioo/personahtmx/models"
)

func Get(username string) *models.User {
    var user models.User;
    initializers.DB.First(&user,"username = ?", username)
    return &user
}
func GetAll()  *[] models.Tweet {
    var Tweets [] models.Tweet;
    initializers.DB.Find(&Tweets)
    return &Tweets
}
