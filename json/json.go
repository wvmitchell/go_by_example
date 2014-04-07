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
  Fruits []string   `json:"frutitas"`
}

func main() {

  // enconding json

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


  res2D := &Response2 {
    Page:  1,
    Fruits: []string{"peach", "pear", "apple", "bananna"}}
  res2B, _ := json.Marshal(res2D)
  fp(string(res2B))


  // Now for decoding

  byt := []byte(`{"num":2.34,"strs":["a","b"]}`)
  var dat map[string]interface{}

  if err := json.Unmarshal(byt, &dat); err != nil {
    panic(err)
  }
  fp(dat)

  // need to cast accessed values to their appropriate type for use
  num := dat["num"].(float64)
  num = num + 1
  fp(num)
}
