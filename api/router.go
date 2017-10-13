package api

import (
  "net/http"
  "log"

  "github.com/bs-online-judge/api/submissions"

  "github.com/gorilla/mux"
)

func Run() {
  r := mux.NewRouter()
  r.HandleFunc("/submissions", Chain(submissions.Get)).Methods("GET")
  http.Handle("/", r)
  log.Fatal(http.ListenAndServe(":8080", r))
}
