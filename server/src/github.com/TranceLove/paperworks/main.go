package main

import (
	// "os"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/TranceLove/paperworks/routes"
	// "github.com/appleboy/gin-jwt"
)

func main(){
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/hello", routes.Hello)
	r.GET("/random-image", routes.RandomImage)

    http.ListenAndServe(":3000", r)
}
