package routes

import (
    "os"
    "fmt"
    "testing"
    "net/http"
    "github.com/TranceLove/paperworks/db"
    "github.com/TranceLove/paperworks/models"
    "github.com/TranceLove/paperworks/routes"
    "github.com/jinzhu/gorm"

    "github.com/stretchr/testify/suite"
    "github.com/appleboy/gofight"
    "github.com/buger/jsonparser"
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

func (suite *UserTestSuite) TestLoginWithEmptyRequest(){
    r := gofight.New()

    r.POST("/login").
		SetDebug(true).
		Run(routes.Init(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			suite.Equal(http.StatusUnsupportedMediaType, r.Code)
    })
}

func (suite *UserTestSuite) TestLoginWithoutProperContentType(){
    r := gofight.New()

    r.POST("/login").
		SetDebug(true).
        SetJSON(gofight.D{}).
		Run(routes.Init(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			suite.Equal(http.StatusUnsupportedMediaType, r.Code)
    })
}

func (suite *UserTestSuite) TestLoginWithEmptyJson(){
    r := gofight.New()

    r.POST("/login").
		SetDebug(true).
        SetHeader(gofight.H{routes.CONTENT_TYPE_HEADER: routes.CONTENT_TYPE_JSON_API}).
        SetJSON(gofight.D{}).
		Run(routes.Init(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			suite.Equal(http.StatusBadRequest, r.Code)
    })
}

func (suite *UserTestSuite) TestLoginAsNewAccount(){
    r := gofight.New()

    r.POST("/login").
		SetDebug(true).
        SetHeader(gofight.H{routes.CONTENT_TYPE_HEADER: routes.CONTENT_TYPE_JSON_API}).
        SetJSON(gofight.D{
            "device_id": "1234567890",
            "device_udid": "abcdefghijkl",
            "os_type": "xersus",
        }).
		Run(routes.Init(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			suite.Equal(http.StatusCreated, r.Code)
            suite.Equal(routes.CONTENT_TYPE_JSON_API, rq.Header.Get(routes.CONTENT_TYPE_HEADER))
            verify, _ := jsonparser.GetString([]byte(r.Body.String()), "data", "type")
            suite.Equal("user", verify)
            id, _ := jsonparser.GetInt([]byte(r.Body.String()), "data", "id")
            suite.True(id > 0)
            deviceUdid, _ := jsonparser.GetString([]byte(r.Body.String()), "data", "attributes", "device_udid")
            osType, _ := jsonparser.GetString([]byte(r.Body.String()), "data", "attributes", "os_type")
            suite.Equal("abcdefghijkl", deviceUdid)
            suite.Equal("xersus", osType)
    })
}

func (suite *UserTestSuite) TestLoginWithCreatedAccount(){
    r := gofight.New()

    ch := make(chan models.User)
    go db.ExecuteGormTemplate(func(engine *gorm.DB) {
        user := models.User {
            DeviceId: "3456789012",
            DeviceUdid: "mnopqrstuvwxyz",
            OsType: "verdetta",
        }

        engine.Create(&user)

        ch <- user
        close(ch)
    })

    user, _ := <- ch

    r.POST("/login").
		SetDebug(true).
        SetHeader(gofight.H{routes.CONTENT_TYPE_HEADER: routes.CONTENT_TYPE_JSON_API}).
        SetJSON(gofight.D{
            "device_id": "3456789012",
            "device_udid": "mnopqrstuvwxyz",
            "os_type": "verdetta",
        }).
		Run(routes.Init(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			suite.Equal(http.StatusOK, r.Code)
            suite.Equal(routes.CONTENT_TYPE_JSON_API, rq.Header.Get(routes.CONTENT_TYPE_HEADER))
            verify, _ := jsonparser.GetString([]byte(r.Body.String()), "data", "type")
            suite.Equal("user", verify)
            id, _ := jsonparser.GetInt([]byte(r.Body.String()), "data", "id")
            suite.True(id > 0)
            suite.Equal(int64(user.ID), id)
            deviceUdid, _ := jsonparser.GetString([]byte(r.Body.String()), "data", "attributes", "device_udid")
            osType, _ := jsonparser.GetString([]byte(r.Body.String()), "data", "attributes", "os_type")
            suite.Equal(user.DeviceUdid, deviceUdid)
            suite.Equal(user.OsType, osType)
    })
}

func TestExampleTestSuite(t *testing.T) {
    suite.Run(t, new(UserTestSuite))
}
