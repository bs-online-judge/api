package main

import (
  "os"
  "log"
  "net/http"

  "github.com/bs-online-judge/api/api"
  "github.com/bs-online-judge/api/redis"
  "github.com/bs-online-judge/api/models"
)

func main() {
  if err := models.CreateTables(); err != nil {
    log.Fatal(err)
  }

  if err := redis.Connect(); err != nil {
    log.Fatal(err)
  }

  log.Fatal(http.ListenAndServe(":" + os.Getenv("API_PORT"), api.Router()))
}
