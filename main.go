package main

import (
  // "net/http"
  // "github.com/mwharris1313/goapi/api"
  "github.com/gin-gonic/gin"
)

func main()  {
  // http.HandleFunc(api.RootPath, api.RootHandler)
  // http.HandleFunc(api.AboutPath, api.AboutHandler)
  // http.ListenAndServe(":3000", nil)

  router := gin.Default()
  router.Static("/", "./public")
  // router.Static("/mp", "./public/mp")
  // router.StaticFS("/more_static", http.Dir("my_file_system"))
  // router.StaticFile("/favicon.ico", "./resources/favicon.ico")

  router.Run(":3000")
}
