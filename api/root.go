package api

import (
  "net/http"
  "fmt"
)

var RootPath = "/"

func RootHandler(w http.ResponseWriter, r *http.Request)  {
  fmt.Fprintf(w, "Root, %s!", r.URL.Path[1:])
}
