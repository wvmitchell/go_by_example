package main

import (
  //"bufio"
  "fmt"
  //"io"
  "io/ioutil"
  //"os"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func main() {
  p := fmt.Println

  dat, err := ioutil.ReadFile("/tmp/dat")
  check(err)
  p(string(dat))
}

