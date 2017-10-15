package api

import "net/http"

func Chain(middlewares ...func(*App) bool) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    app, err := NewApp(w, r)
    if err != nil {
      return
    }
    for _, m := range middlewares {
      if !m(app) {
        return
      }
    }
  }
}
