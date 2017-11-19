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

func NewApp(w http.ResponseWriter, r *http.Request) *App {
  sessionCookie, err := r.Cookie("sessionId")

  // no session cookie
  if err != nil {
    return &App{response: w, request: r}
  }

  // fetch user session
  var userSession *redis.UserSession
  userSession, err = redis.GetUserSession(sessionCookie.Value)

  // invalid session
  if err != nil {
    return &App{response: w, request: r}
  }

  return &App{w, r, userSession}
}
