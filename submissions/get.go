package submissions

import (
  "fmt"
  "net/http"
)

func Get(w http.ResponseWriter, r *http.Request) bool {
  fmt.Println("GET /submissions")
  fmt.Fprintf(w, "GET /submissions")
  return true
}
