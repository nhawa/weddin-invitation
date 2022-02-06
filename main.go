package main

import (
    "strings"

    log "github.com/sirupsen/logrus"
    "github.com/spf13/viper"
)

func init() {
    viper.AutomaticEnv()
    viper.SetConfigType("yaml")

    replacer := strings.NewReplacer(".", "_")
    viper.SetEnvKeyReplacer(replacer)

    // Set default
    viper.SetDefault("application.port", 9010)

    viper.SetConfigFile(`config.yml`)
    err := viper.ReadInConfig()
    if err != nil {
        log.Error(err)
    }
}

func main() {
    log.SetFormatter(&log.JSONFormatter{})

    app := NewApplication()

    app.ListenAndServe()
}
