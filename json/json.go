package main

import (
  "encoding/json"
  "fmt"
  //"os"
)

var fp = fmt.Println

type Response1 struct {
  Page int
  Fruits []string
}

type Response2 struct {
  Page int          `json:"page"`
  Fruits []string   `json:"fruits"`
}

func main() {

  bolB, _ := json.Marshal(true)
  fp(string(bolB))

  intB, _ := json.Marshal(1)
  fp(string(intB))

  fltB, _ := json.Marshal(12.334)
  fp(string(fltB))

  strB, _ := json.Marshal("not a beaver, but a gopher")
  fp(string(strB))

  mapD := map[string]int{"apple": 1, "peach": 4}
  mapB, _ := json.Marshal(mapD)
  fp(string(mapB))

  res1D := &Response1 {
    Page: 1,
    Fruits:  []string{"peach", "pear", "apple", "bananna"}}
  res1B, _ := json.Marshal(res1D)
  fp(string(res1B))
}
