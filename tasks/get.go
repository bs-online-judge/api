package tasks

import (
  "fmt"
  "net/http"
)

func GetAll(w http.ResponseWriter, r *http.Request) bool {
  fmt.Println("GET /queue")
  fmt.Fprintf(w, "GET /queue")
  return true
}
