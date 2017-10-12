package models

import (
    "os"
    "fmt"
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/TranceLove/paperworks/config"
)

func TestConfigFileNotFound(t *testing.T){
    _, ok := models.GetConfigurationFor("dev")
    assert.False(t, ok)
}

func TestReadConfigFile(t *testing.T){
    cwd, _ := os.Getwd()
    os.Setenv("PAPERWORKS_CONFIG", fmt.Sprintf("%s/config1.json", cwd))
    result, ok := models.GetConfigurationFor("dev")
    assert.True(t, ok)
    assert.NotNil(t, result)
    assert.Equal(t, "localhost", result.DatabaseConfig.Host)
    assert.Equal(t, 5432, result.DatabaseConfig.Port)
    assert.Equal(t, "paperworks", result.DatabaseConfig.DatabaseName)
    assert.Equal(t, "postgres", result.DatabaseConfig.Username)
    assert.Equal(t, "postgres", result.DatabaseConfig.Password)
}

func TestReadConfigFile2(t *testing.T){
    cwd, _ := os.Getwd()
    os.Setenv("PAPERWORKS_CONFIG", fmt.Sprintf("%s/config2.json", cwd))
    result, ok := models.GetConfigurationFor("dev")
    assert.True(t, ok)
    assert.NotNil(t, result)
    assert.Equal(t, "localhost", result.DatabaseConfig.Host)
    assert.Equal(t, 5432, result.DatabaseConfig.Port)
    assert.Equal(t, "paperworks", result.DatabaseConfig.DatabaseName)
    assert.Equal(t, "postgres", result.DatabaseConfig.Username)
    assert.Equal(t, "postgres", result.DatabaseConfig.Password)

    result, ok = models.GetConfigurationFor("staging")
    assert.True(t, ok)
    assert.NotNil(t, result)
    assert.Equal(t, "12-34-56-78.pgsql.rds.amazonaws.com", result.DatabaseConfig.Host)
    assert.Equal(t, 25432, result.DatabaseConfig.Port)
    assert.Equal(t, "paperworks-staging", result.DatabaseConfig.DatabaseName)
    assert.Equal(t, "postadmina", result.DatabaseConfig.Username)
    assert.Equal(t, "w3s7G0", result.DatabaseConfig.Password)

    result, ok = models.GetConfigurationFor("production")
    assert.True(t, ok)
    assert.NotNil(t, result)
    assert.Equal(t, "12-34-56-78.pgsql.rds.amazonaws.com", result.DatabaseConfig.Host)
    assert.Equal(t, 55432, result.DatabaseConfig.Port)
    assert.Equal(t, "paperworks-production", result.DatabaseConfig.DatabaseName)
    assert.Equal(t, "postadmina", result.DatabaseConfig.Username)
    assert.Equal(t, "F7ju6sU#==", result.DatabaseConfig.Password)
}
