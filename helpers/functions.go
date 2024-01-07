package helpers

import (
	"errors"

	"github.com/MarcosIgnacioo/personahtmx/crud"
	"github.com/MarcosIgnacioo/personahtmx/models"
	"github.com/gin-contrib/sessions"
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

func CreateSession(user models.User, c *gin.Context){
    sessionStorage := sessions.Default(c)
    sessionStorage.Set("Username", user.Username)
}
func GetSession(c *gin.Context) (*models.User, error) {
    sessionStorage := sessions.Default(c)

    usernameValue := sessionStorage.Get("Username")

    if usernameValue == nil {
        return nil, errors.New("Not logged in")
    }

    username, ok := usernameValue.(string)
    if !ok {
        return nil, errors.New("Failed to assert Username as string")
    }

    user := crud.Get(username)
    return user, nil
}
