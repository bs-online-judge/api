package api

import (
  "net/http"
  "encoding/json"

  "github.com/bs-online-judge/api/redis"
  "github.com/bs-online-judge/api/models"
  "github.com/bs-online-judge/api/helpers"
)

type PostUserData struct {
  Username string
  Password string
}

func PostUser(app *App) bool {
  decoder := json.NewDecoder(app.request.Body)
  data := PostUserData{}
  if err := decoder.Decode(&data); err != nil {
    app.response.WriteHeader(http.StatusBadRequest)
    return false
  }
  if err := models.CreateUser(data.Username, data.Password); err != nil {
    app.response.WriteHeader(http.StatusInternalServerError)
    return false
  }
  return true
}

func Login(app *App) bool {
  username, password, ok := app.request.BasicAuth()
  if !ok {
    helpers.Unauthorized(app.response)
    return false
  }

  user, err := models.GetUserByCredentials(username, password)
  if err != nil {
    helpers.Unauthorized(app.response)
    return false
  }

  var key string
  key, err = helpers.NewUUID()
  if err != nil {
    app.response.WriteHeader(http.StatusInternalServerError)
    return false
  }
  key += username

  if err := redis.SetUserSession(key, user); err != nil {
    app.response.WriteHeader(http.StatusInternalServerError)
    return false
  }

  cookie := http.Cookie{Name: "sessionId", Value: key}
  http.SetCookie(app.response, &cookie)
  app.response.WriteHeader(http.StatusOK)

  return true
}
