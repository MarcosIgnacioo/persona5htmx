package helpers

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/MarcosIgnacioo/personahtmx/crud"
	"github.com/MarcosIgnacioo/personahtmx/initializers"
	"github.com/MarcosIgnacioo/personahtmx/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func CreateSession(user *models.User, c *gin.Context) {
    session, _ := initializers.Store.Get(c.Request, "session")
    session.Values["user"] = user
    session.Save(c.Request, c.Writer)
}
func GetUserSession(c *gin.Context) (*models.User, error) {
    session, _ := initializers.Store.Get(c.Request, "session")
    var user = &models.User{}
    val := session.Values["user"]
    var ok bool
    if user, ok = val.(*models.User); !ok {
        return nil, errors.New("Error at GetSession")
    }
    return user, nil
}

func LoadTweets(c *gin.Context )  models.Tweets {
    tweetLoad := c.Query("load")
    load, error := strconv.Atoi(tweetLoad)
    if error != nil {
        load = 0
    }
    template := "tweets.html"
    if load == 0 {
        template = "index.html" 
    }
    fmt.Println("////")
    fmt.Println(load)
    fmt.Println(tweetLoad)
    fmt.Println("\\\\\\")
    tweets := crud.GetAll(load)



    // TODO hacer un join de las columnas de contenido y autor de la tabla replies con el id de LastReplyID para cada registro y agregarle al struct de abajo una nueva propiedad que sea replies que sea un arreglo de tweets igualmente 

    return models.Tweets { Tweets: tweets, Start: load, Next: load + 2,More: load < 4, Template: template}

}
