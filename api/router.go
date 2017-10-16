package api

import "github.com/gorilla/mux"

func Router() *mux.Router {
  r := mux.NewRouter()
  r.HandleFunc("/tasks", Chain(GetAllTasks)).Methods("GET")
  r.HandleFunc("/user", Chain(PostUser)).Methods("POST")
  return r
}
