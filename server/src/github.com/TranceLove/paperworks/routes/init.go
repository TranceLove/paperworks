package routes

import (
    "github.com/gin-gonic/gin"
)

const (
    CONTENT_TYPE_HEADER string = "Content-Type"
    CONTENT_TYPE_JSON_API string = "application/vnd.api+json"
)

func Init() *gin.Engine {
    r := gin.New()
    r.Use(gin.Logger())
    r.Use(gin.Recovery())

    r.GET("/hello", hello)
	r.GET("/random-image", randomImage)

    r.POST("/login", loginOrCreateAccount)

    return r
}
