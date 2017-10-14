package api

import (
  "github.com/bs-online-judge/api/queue"

  "github.com/gorilla/mux"
)

func Router() *mux.Router {
  r := mux.NewRouter()
  r.HandleFunc("/queue", Chain(queue.Get)).Methods("GET")
  return r
}
