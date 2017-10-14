package models

import (
    "github.com/jinzhu/gorm"
)

type User struct {
    gorm.Model
    Username string `gorm:"size:255;not null;unique"`
    DeviceId string `gorm:"size:256;not null;column_name:'device_id'"`
    OsType string `gorm:"size:32;not null"`
    Role string `gorm:"size:16;not null;default:''"`
}
