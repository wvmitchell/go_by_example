package main

import (
  "os"
  "fmt"
  //"strings"
)

func main() {

  var p = fmt.Println

  os.Setenv("FOO", "1")
  p("FOO", os.Getenv("FOO"))
  p("BAR", os.Getenv("BAR"))
}
