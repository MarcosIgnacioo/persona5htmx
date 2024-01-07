package models

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
    gorm.Model
    Username string `gorm:"column:username;unique"`
    Password string `gorm:"password"`
}

func (u User) String() string  {
    return fmt.Sprintf("Name: %s \nPassword: %s \n", u.Username, u.Password)
}
