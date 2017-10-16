package api

import (
  "net/http"
  "encoding/json"
  
  "github.com/bs-online-judge/api/models"
)

type PostUserData struct {
  Username string
  Password string
}

func PostUser(app *App) bool {
  if app.user != nil {
    app.response.WriteHeader(http.StatusForbidden)
    return false
  }
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
