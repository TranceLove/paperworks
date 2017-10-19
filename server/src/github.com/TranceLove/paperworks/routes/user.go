package routes

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    "github.com/TranceLove/paperworks/db"
    "github.com/TranceLove/paperworks/models"
)

type loginForm struct {
    DeviceId string `json:"device_id" binding:"required"`
    DeviceUdid string `json:"device_udid" binding:"required"`
    OsType string `json:"os_type"`
}

func _sendLoginResponse(c *gin.Context, status int, user models.User) {
    setJsonApiResponseContentType(c)
    c.JSON(status, gin.H{
        "data": gin.H{
            "type": "user",
            "id": user.ID,
            "attributes": gin.H{
                "device_udid": user.DeviceUdid,
                "os_type": user.OsType,
            },
        },
    })
}

func loginOrCreateAccount(c *gin.Context){
    verifyRequestValidity(c)

    var request loginForm
    c.BindJSON(&request)
    if(request.DeviceId == "" || request.DeviceUdid == ""){
        c.AbortWithStatus(http.StatusBadRequest)
    } else {
        db.ExecuteGormTemplate(func(engine *gorm.DB) {
            var user models.User
            engine.Where(&models.User{DeviceId:request.DeviceId,DeviceUdid:request.DeviceUdid}).First(&user)
            if(user.ID > 0){
                _sendLoginResponse(c, http.StatusOK, user)
            } else {
                tx := engine.Begin()
                user = models.User{
                    DeviceId:request.DeviceId,
                    DeviceUdid:request.DeviceUdid,
                    OsType:request.OsType,
                }
                err := engine.Create(&user).Error
                if(err != nil){
                    tx.Rollback()
                    c.AbortWithStatus(503)
                } else {
                    _sendLoginResponse(c, http.StatusCreated, user)
                }
            }
        })
    }
}
