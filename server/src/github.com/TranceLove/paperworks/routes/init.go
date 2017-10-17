package routes

import (
    "github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
    r := gin.New()
    r.Use(gin.Logger())
    r.Use(gin.Recovery())

    r.GET("/hello", hello)
	r.GET("/random-image", randomImage)

    return r
}
