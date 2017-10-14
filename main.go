package main

import (
  "net/http"
  "log"

  "github.com/bs-online-judge/api/api"
)

func main() {
  log.Fatal(http.ListenAndServe(":8080", api.Router()))
}
