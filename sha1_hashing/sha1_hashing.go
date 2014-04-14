package main

import (
  "crypto/sha1"
  "crypto/md5"
  "fmt"
)

func main() {

  p := fmt.Println

  s := "Hash this string"

  h := sha1.New()

  h.Write([]byte(s))

  bs := h.Sum(nil)

  p(s)
  fmt.Printf("%x\n", bs)


  // same pattern for MD5 hashes
  hs := md5.New()
  hs.Write([]byte(s))
  ms := hs.Sum(nil)
  fmt.Printf("%x\n", ms)

}
