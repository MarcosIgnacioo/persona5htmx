package helpers

import (
	"errors"

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
