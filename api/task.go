package api

import (
  "fmt"

  "github.com/bs-online-judge/api/models"
)

func GetAllTasks(app *App) bool {
  t := models.Task{}
  fmt.Println("GET /queue", t.Id)
  fmt.Fprintf(app.response, "GET /queueeeee")
  return true
}
