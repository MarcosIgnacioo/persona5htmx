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
func GetAll(offset int)  [] models.Tweet {
    var Tweets [] models.Tweet;
    initializers.DB.Limit(2).Offset(offset).Find(&Tweets)
    return Tweets
}
