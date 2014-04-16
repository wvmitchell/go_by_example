package main

import (
  //"bufio"
  "fmt"
  "io/ioutil"
  "os"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func main() {

  // dump string into a file
  d1 := []byte("hello\nhere is some\ngo")
  err := ioutil.WriteFile("/tmp/dat1", d1, 0644)
  check(err)

  // more granular writes
  f, err := os.Create("/tmp/dat2")
  check(err)

  // defer closing the file immediately after opening
  defer f.Close()

  // writing bytes to a file
  d2 := []byte{115, 111, 109, 101, 10}
  n2, err := f.Write(d2)
  fmt.Printf("Wrote %d bytes\n", n2)



}
