package api

import (
  "net/http"
  "fmt"
)

var AboutPath = "/about/"

func AboutHandler(w http.ResponseWriter, r *http.Request)  {
  fmt.Fprintf(w, "About, %s!", r.URL.Path[1:])
}
