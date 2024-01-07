package models

import "gorm.io/gorm"

type Tweet struct {
    gorm.Model
    Author string `gorm:"column:author"`
    Content string `gorm:"column:content"`
}
