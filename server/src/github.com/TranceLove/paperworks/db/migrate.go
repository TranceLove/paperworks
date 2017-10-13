package db

import (
    "fmt"
    "log"
    "github.com/jinzhu/gorm"
    "github.com/TranceLove/paperworks/models"
)

func Migrate(){
    ch := make(chan interface{})
    go ExecuteGormTemplate(func(engine *gorm.DB) {
        fmt.Println("Start schema migration")
        engine.AutoMigrate(&models.User{})
        ch <- "OK"
        close(ch)
    })
    _, ok := <- ch
    if(!ok){
        log.Fatal("Error performing schema migration")
    } else {
        fmt.Println("Schema migration completed")
    }
}
