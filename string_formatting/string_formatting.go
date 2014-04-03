package main

import (
  "fmt"
  //"os"
)

var pf = fmt.Printf

type point struct {
  x, y int
}

func main(){
  p := point{1,3}

  // => {1, 3}
  pf("%v\n", p)

  // => {x:1 y:3}
  pf("%+v\n", p)

  // => main.point{x:1, y:3}
  pf("%#v\n", p)
}
