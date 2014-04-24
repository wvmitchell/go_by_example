package main

import (
  "os"
  "fmt"
  //"strings"
)

func main() {

  var p = fmt.Prinln

  os.Setenv("FOO", "1")
  p("FOO", os.Getenv("FOO")
}
