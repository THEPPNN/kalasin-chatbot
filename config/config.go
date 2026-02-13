package config

import (
    "log"
    "os"
    "github.com/joho/godotenv"
)

func Init() {
    wd, _ := os.Getwd()
    log.Println("WORKDIR:", wd)

    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env")
    }
}

func Get(key string) string {
    return os.Getenv(key)
}