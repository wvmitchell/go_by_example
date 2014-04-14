package main

import (
  b64 "encoding/base64"
  "fmt"
)

func main() {

  // encoding and decoding a string
  p := fmt.Println
  data := "123abc!?$*&()'-=@~"

  sEnc := b64.StdEncoding.EncodeToString([]byte(data))
  p(sEnc)

  dEnc, _ := b64.StdEncoding.DecodeString(sEnc)
  p(string(dEnc))

  // URL compatible base64 encoding and decoding
  uEnc := b64.URLEncoding.EncodeToString([]byte(data))
  p(uEnc)

  duEnc, _ := b64.URLEncoding.DecodeString(uEnc)
  p(string(duEnc))

}
