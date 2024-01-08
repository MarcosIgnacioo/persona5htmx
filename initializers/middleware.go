package initializers

import (
	"encoding/gob"

	"github.com/MarcosIgnacioo/personahtmx/models"
	"github.com/gorilla/sessions"
)

var Store = sessions.NewCookieStore([]byte("super-secret"))

func SetupCookies()  {
	gob.Register(&models.User{})
}
