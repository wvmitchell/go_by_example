package main

import (
  //"bufio"
  "fmt"
  //"io"
  "io/ioutil"
  "os"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func main() {
  p := fmt.Println

  // read file
  dat, err := ioutil.ReadFile("/tmp/dat")
  check(err)
  p(string(dat))


  // open file
  file, err := os.Open("/tmp/dat")

  b1 := make([]byte, 5)
  n1, err := file.Read(b1)
  check(err)
  fmt.Printf("%d bytes: %s\n", n1, string(b1))
}

