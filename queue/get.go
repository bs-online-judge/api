package queue

import (
  "fmt"
  "net/http"
)

func Get(w http.ResponseWriter, r *http.Request) bool {
  fmt.Println("GET /queue")
  fmt.Fprintf(w, "GET /queue")
  return true
}
