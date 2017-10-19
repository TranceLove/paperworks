package routes

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func verifyRequestValidity(c *gin.Context){
    if(c.GetHeader(CONTENT_TYPE_HEADER) != CONTENT_TYPE_JSON_API){
        c.AbortWithStatus(http.StatusUnsupportedMediaType)
        return
    }
}

func setJsonApiResponseContentType(c *gin.Context){
    c.Header(CONTENT_TYPE_HEADER, CONTENT_TYPE_JSON_API)
}
