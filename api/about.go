package api

import (
  "net/http"
  "encoding/json"
)

type Message struct{
  Message string `json:"msg"`
}

type Wrapper struct{
  WMessage string `json:"wmsg"`
  Data Message `json:"data"`
}

var AboutPath = "/about/"

func AboutHandler(w http.ResponseWriter, r *http.Request)  {
  msg := Message{"goapi API v0a (alpha)"}
  wrap := Wrapper{
    WMessage: "WrapperMessage",
    Data: msg,
  }
  // jsonMsg, err := json.Marshal(msg)
  //
  // if err != nil {
  //   panic(err)
  // }

  jsonWrap, err := json.Marshal(wrap)
  if err != nil {
    panic(err)
  }

  w.Write(jsonWrap)
}
