package routes

import (
    "os"
    "fmt"
    "net/http"
    "testing"

    "github.com/TranceLove/paperworks/routes"

    "github.com/buger/jsonparser"

    "github.com/appleboy/gofight"
    "github.com/stretchr/testify/suite"
)

type HelloTestSuite struct {
    suite.Suite
}

func (suite *HelloTestSuite) SetupSuite(){
    cwd, _ := os.Getwd()
    os.Setenv("PAPERWORKS_CONFIG", fmt.Sprintf("%s/../test_config.json", cwd))
    os.Setenv("PAPERWORKS_ENV", "test")
}

func (suite *HelloTestSuite) TestRoot(){
    r := gofight.New()

    r.GET("/").
		SetDebug(true).
		Run(routes.Init(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			suite.Equal(http.StatusNotFound, r.Code)
    })
}

func (suite *HelloTestSuite) TestHello(){
    r := gofight.New()

    r.GET("/hello").
		SetDebug(true).
        SetHeader(gofight.H{
            "Content-Type": routes.CONTENT_TYPE_JSON_API,
        }).
		Run(routes.Init(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			suite.Equal(http.StatusOK, r.Code)
            verify, _ := jsonparser.GetString([]byte(r.Body.String()), "data", "message")
            suite.Equal("Hello World", verify)
    })
}

func TestExampleTestSuite(t *testing.T) {
    suite.Run(t, new(HelloTestSuite))
}
