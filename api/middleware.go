package api

import "net/http"

func Chain(middlewares ...func(*App) bool) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    app := NewApp(w, r)
    for _, m := range middlewares {
      if !m(app) {
        return
      }
    }
  }
}
