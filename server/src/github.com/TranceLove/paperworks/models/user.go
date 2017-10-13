package models

import (
    "github.com/jinzhu/gorm"
)

type User struct {
    gorm.Model
    Username string `gorm:"size:255"`
    Device_id string `gorm:"size:256"`
}
