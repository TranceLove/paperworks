package routes

import (
    "time"
    "github.com/gin-gonic/gin"
)

func hello(c *gin.Context) {
    setJsonApiResponseContentType(c)
    c.JSON(200, gin.H{
        "data": gin.H{
            "message": "Hello World",
            "time": time.Now(),
        },
    })
}
