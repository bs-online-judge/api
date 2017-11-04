package api

func HealthCheck(app *App) bool {
  app.response.Write([]byte("WORKING"))

  return true
}
