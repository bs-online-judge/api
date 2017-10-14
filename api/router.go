package api

import (
  "github.com/bs-online-judge/api/submissions"

  "github.com/gorilla/mux"
)

func Router() *mux.Router {
  r := mux.NewRouter()
  r.HandleFunc("/submissions", Chain(submissions.Get)).Methods("GET")
  return r
}
