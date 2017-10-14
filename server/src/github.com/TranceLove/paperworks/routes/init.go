package routes

import (
    "github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
    r.GET("/hello", hello)
	r.GET("/random-image", randomImage)
}
