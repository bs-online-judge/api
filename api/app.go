package api

import (
  "net/http"

  "github.com/bs-online-judge/api/redis"
)

type App struct {
  response http.ResponseWriter
  request *http.Request
  userSession *redis.UserSession
}

func NewApp(w http.ResponseWriter, r *http.Request) (*App, error) {
  sessionCookie, err := r.Cookie("sessionId")


  // User not logged
  if err != nil {
    return &App{response: w, request: r}, nil
  }

  var userSession *redis.UserSession
  userSession, err = redis.GetUserSession(sessionCookie.Value)

  // Session exists
  if err != nil {
    return &App{response: w, request: r}, nil
  }

  return &App{w, r, userSession}, nil
}
