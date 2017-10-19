package models

import (
    "os"
    "fmt"
    "testing"
    "github.com/TranceLove/paperworks/db"
    "github.com/TranceLove/paperworks/models"
    "github.com/jinzhu/gorm"

    "github.com/stretchr/testify/suite"
)

type UserTestSuite struct {
    suite.Suite
}

func (suite *UserTestSuite) SetupSuite(){
    cwd, _ := os.Getwd()
    os.Setenv("PAPERWORKS_CONFIG", fmt.Sprintf("%s/../test_config.json", cwd))
    os.Setenv("PAPERWORKS_ENV", "test")

    db.ExecuteGormTemplate(func(engine *gorm.DB) {
        engine.CreateTable(&models.User{})
    })
}

func (suite *UserTestSuite) TearDownSuite(){
    db.ExecuteGormTemplate(func(engine *gorm.DB) {
        engine.DropTable(&models.User{})
    })
}

func (suite *UserTestSuite) TestOnePlusOne(){
    suite.Equal(2, 1+1)
}

func (suite *UserTestSuite) TestCreateAndFindUser(){

    db.ExecuteGormTemplate(func(engine *gorm.DB) {
        user := models.User {
            DeviceId: "1234567890",
            DeviceUdid: "abcdefghijkl",
            Username: "TranceLove",
            OsType: "android",
        }

        engine.Create(&user)

        var result models.User
        engine.Where(&models.User{DeviceId:"1234567890"}).First(&result)
        suite.NotNil(result)
        suite.Equal(user.DeviceId, result.DeviceId)
        suite.Equal(user.DeviceUdid, result.DeviceUdid)
        suite.Equal(user.Username, result.Username)
        suite.Equal(user.OsType, result.OsType)

        id := result.ID

        engine.Where("id = ?", id).First(&result)
        suite.Equal(user.DeviceId, result.DeviceId)
        suite.Equal(user.DeviceUdid, result.DeviceUdid)
        suite.Equal(user.Username, result.Username)
        suite.Equal(user.OsType, result.OsType)
    })
}

func TestExampleTestSuite(t *testing.T) {
    suite.Run(t, new(UserTestSuite))
}
