package api

import (
  "net/http"
  "encoding/json"
)

type Message struct{
  Text string
}

var AboutPath = "/about/"

func AboutHandler(w http.ResponseWriter, r *http.Request)  {
  msg := Message{"goapi API v0a (alpha)"}
  jsonMsg, err := json.Marshal(msg)

  if err != nil {
    panic(err)
  }

  w.Write(jsonMsg)
}
