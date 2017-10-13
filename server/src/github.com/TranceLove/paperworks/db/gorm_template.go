package db

import (
    "os"
    "fmt"
    "log"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "github.com/TranceLove/paperworks/config"
)

type gormTemplate func(engine *gorm.DB)

const (
    GORM_DIALECT = "postgres"
    GORM_CONNECTION_URL = "user=%s password=%s host=%s port=%d dbname=%s sslmode=disable"
)

func ExecuteGormTemplate(fn gormTemplate) {
    env := os.Getenv("PAPERWORKS_ENV")
    if(env == ""){
        env = "dev"
    }
    var err error
    config, _ := config.GetConfigurationFor(env)
    dbConfig := config.DatabaseConfig
    engine, err := gorm.Open(GORM_DIALECT, fmt.Sprintf(GORM_CONNECTION_URL, dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DatabaseName))
    if(err != nil){
        log.Fatal("Failed to initialize gorm.", err)
    }

    defer engine.Close()

    fn(engine)
}
