package router

import (
  "net/http"
)

// middleware to protect private pages
func corsHandler(h http.Handler) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Content-Type", "application/json; charset=utf-8")
    w.Header().Add("Access-Control-Allow-Origin", "*")
    w.Header().Add("Access-Control-Allow-Credentials", "true")
    w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    w.Header().Add("Access-Control-Allow-Headers",
        "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")

    h.ServeHTTP(w,r)
  }
}

func handleOptions(w http.ResponseWriter, r *http.Request) {
  responseText(w, "Done", http.StatusOK)
}
