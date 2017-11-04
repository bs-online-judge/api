package api

import "github.com/gorilla/mux"

func Router() *mux.Router {
  r := mux.NewRouter()
  r.HandleFunc("/tasks", Chain(GetAllTasks, PostUser)).Methods("GET")
  r.HandleFunc("/user", Chain(PostUser)).Methods("POST")
  r.HandleFunc("/user/login", Chain(Login)).Methods("POST")
  r.HandleFunc("/healthcheck", Chain(HealthCheck)).Methods("GET")
  return r
}
