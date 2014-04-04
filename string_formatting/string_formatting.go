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

  // => main.point
  pf("%T\n", p)

  // => true
  pf("%t\n", true)

  // Formatting integers
  // => 123
  pf("%d\n", 123)

  // => 1111011
  pf("%b\n", 123)

  // Character corresponding to int
  // => v
  pf("%c\n", 118)

  // Hex encoding
  // => 32b4ee
  pf("%x\n", 3323118)

  // Floats
  // => 13.300000
  pf("%f\n", 13.3)

  // Scientific notation
  // => 1.230000e+10
  pf("%e\n", 12300000000.0)

  // => 1.230000E+10
  pf("%E\n", 12300000000.0)

  // => "string"
  pf("%s\n", "\"string\"")

  // => "\"string\""
  pf("%q\n", "\"string\"")

  // => 686578
  pf("%x\n", "hex")

}
