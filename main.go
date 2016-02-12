package main

import (
  "net/http"
  "github.com/mwharris1313/goapi/api"
  "github.com/gin-gonic/gin"
)

func main()  {
  r := gin.Default()

  http.HandleFunc(api.RootPath, api.RootHandler)
  http.HandleFunc(api.AboutPath, api.AboutHandler)

  http.ListenAndServe(":3000", nil)
  r.Run(":3000")
}
