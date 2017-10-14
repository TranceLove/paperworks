package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/TranceLove/paperworks/routes"
	// "github.com/appleboy/gin-jwt"
)

func main(){

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	routes.Init(r)

    http.ListenAndServe(":3000", r)
}
