package helpers

import "net/http"

func Unauthorized(w http.ResponseWriter) {
  w.Header().Set("WWW-Authenticate", "Basic")
  w.WriteHeader(http.StatusUnauthorized)
}
