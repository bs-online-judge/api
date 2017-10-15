package api

import (
  "github.com/bs-online-judge/api/tasks"

  "github.com/gorilla/mux"
)

func Router() *mux.Router {
  r := mux.NewRouter()
  r.HandleFunc("/tasks", Chain(tasks.GetAll)).Methods("GET")
  return r
}
