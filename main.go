package main

import (
  "net/http"
  "github.com/mwharris1313/goapi/api"
)

func main()  {
  http.HandleFunc(api.RootPath, api.RootHandler)
  http.HandleFunc(api.AboutPath, api.AboutHandler)

  http.ListenAndServe(":3000", nil)
}
