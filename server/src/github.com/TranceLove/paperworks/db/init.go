package db

import (
    "os"
    "fmt"
    "log"
    "github.com/go-xorm/xorm"
    _ "github.com/lib/pq"
    "github.com/TranceLove/paperworks/config"
)

var engine *xorm.Engine

func Init(){
    env := os.Getenv("PAPERWORKS_ENV")
    if(env == ""){
        env = "dev"
    }
    var err error
    config, _ := config.GetConfigurationFor(env)
    dbConfig := config.DatabaseConfig
    engine, err = xorm.NewEngine("postgres", fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DatabaseName))
    if(err != nil){
        log.Fatal("Failed to initialize xorm.", err)
    }
}

func GetXormEngine() *xorm.Engine {
    return engine
}
