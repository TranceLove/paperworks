package db

import (
    "fmt"
    "github.com/jinzhu/gorm"
)

func Migrate(){
    ch := make(chan interface{})
    go ExecuteGormTemplate(func(engine *gorm.DB) {
        fmt.Println("Hello")
        ch <- 1
        close(ch)
    })
    result, ok := <- ch
    if(!ok){
        fmt.Println("Error")
    } else {
        fmt.Println(result)
    }
}
