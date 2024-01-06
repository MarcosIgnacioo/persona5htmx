package models

import "gorm.io/gorm"

type Prueba struct {
  gorm.Model
    Name string `gorm:"column:name"`
}
