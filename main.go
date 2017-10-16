package main

import (
  "os"
  "log"
  "net/http"
  
  "github.com/bs-online-judge/api/models"
  "github.com/bs-online-judge/api/api"
)

func main() {
  if err := models.CreateTables(); err != nil {
    log.Fatal(err)
  }
  log.Fatal(http.ListenAndServe(":" + os.Getenv("API_PORT"), api.Router()))
}
