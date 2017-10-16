package api

import (
  "net/http"

  "github.com/bs-online-judge/api/models"
  "github.com/bs-online-judge/api/helpers"
)

type App struct {
  response http.ResponseWriter
  request *http.Request
  user *models.User
}

func NewApp(w http.ResponseWriter, r *http.Request) (*App, error) {
  username, password, ok := r.BasicAuth()
  if !ok {
    return &App{response: w, request: r}, nil
  }
  user, err := models.GetUserByCredentials(username, password)
  if err != nil {
    helpers.Unauthorized(w)
    return nil, err
  }
  return &App{w, r, user}, nil
}
