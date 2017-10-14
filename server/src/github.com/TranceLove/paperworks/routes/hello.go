package routes

import (
    "time"
    "github.com/gin-gonic/gin"
)

func hello(c *gin.Context) {
    c.JSON(200, gin.H{
        "message": "Hello World",
        "time": time.Now(),
    })
}
