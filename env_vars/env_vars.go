package main

import (
  "os"
  "fmt"
  "strings"
)

func main() {

  var p = fmt.Println

  os.Setenv("FOO", "1")
  p("FOO", os.Getenv("FOO"))

  // can use export BAR=5 prior to running program
  p("BAR", os.Getenv("BAR"))
  p()

  for _, env := range os.Environ() {
    pair := strings.Split(env, "=")
    p(pair[0])
  }
}
