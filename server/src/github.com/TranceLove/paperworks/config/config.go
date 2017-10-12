package models

import (
    "os"
    "fmt"
    "io/ioutil"
    "encoding/json"
)

type config struct {
    DatabaseConfig struct {
        Host string `json:"host"`
        Port int `json:"port"`
        DatabaseName string `json:"databaseName"`
        Username string `json:"username"`
        Password string `json:"password"`
    } `json:"databaseConfig"`
}

func readEnvironmentConfig() map[string]config {
    configFile := os.Getenv("PAPERWORKS_CONFIG")
    if(configFile == "") {
        configFile = "/etc/paperworks/config.json"
    }

    var objmap map[string]*json.RawMessage
    raw, err := ioutil.ReadFile(configFile)
    if err != nil {
        fmt.Println("Error reading configuration file:", err)
        objmap = nil
        return nil
    } else {
        json.Unmarshal(raw, &objmap)

        configs := make(map[string]config)

        for key, value := range objmap {
            var _config config
            json.Unmarshal(*value, &_config)
            configs[key] = _config
        }

        objmap = nil
        raw = nil

        return configs
    }
}

func GetConfigurationFor(env string) (config, bool) {
    retval, ok := readEnvironmentConfig()[env]
    return retval, ok
}
