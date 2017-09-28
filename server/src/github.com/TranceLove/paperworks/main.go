package main

import (
	"github.com/gin-gonic/gin"
	"github.com/appleboy/gin-jwt"
)

func main(){
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

    r.Run(":8080")
}
