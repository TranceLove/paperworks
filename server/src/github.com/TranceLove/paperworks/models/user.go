package models

import (
    "github.com/jinzhu/gorm"
)

type User struct {
    gorm.Model
    Username string `gorm:"size:255;not null;default:'Anonymous User'"`
    DeviceUdid string `gorm:"size:255;not null"`
    DeviceId string `gorm:"size:256;not null;column_name:'device_id'"`
    OsType string `gorm:"size:32;not null;default:'Unknown'"`
    Role string `gorm:"size:16;not null;default:''"`
}

func (u *User) TableName() string {
    return "users"
}
