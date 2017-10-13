package api

import "net/http"

func Chain(middlewares ...func(http.ResponseWriter, *http.Request) bool) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    for _, m := range middlewares {
      if !m(w, r) {
        return
      }
    }
  }
}
