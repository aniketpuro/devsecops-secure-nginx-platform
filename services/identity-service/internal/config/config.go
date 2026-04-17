package config

import (
    "log"
    "time"

    "github.com/spf13/viper"
)

type Config struct {
    Port             string
    Environment      string
    JWTSecret        string
    JWTAccessExpiry  time.Duration
    JWTRefreshExpiry time.Duration
}

func Load() *Config {
    viper.SetConfigName(".env")
    viper.SetConfigType("env")
    viper.AddConfigPath(".")

    viper.AutomaticEnv()

    if err := viper.ReadInConfig(); err != nil {
        log.Println("Warning: No .env file found, using environment variables")
    }

    return &Config{
        Port:             viper.GetString("PORT"),
        Environment:      viper.GetString("ENVIRONMENT"),
        JWTSecret:        viper.GetString("JWT_SECRET"),
        JWTAccessExpiry:  viper.GetDuration("JWT_ACCESS_EXPIRY"),
        JWTRefreshExpiry: viper.GetDuration("JWT_REFRESH_EXPIRY"),
    }
}