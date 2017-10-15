package api

import (
  "net/http"

  "github.com/bs-online-judge/api/models"
)

type App struct {
  response http.ResponseWriter
  request *http.Request
  user *models.User
}

func NewApp(w http.ResponseWriter, r *http.Request) (*App, error) {
  app := &App{
    response: w,
    request: r,
  }
  return app, nil
}
