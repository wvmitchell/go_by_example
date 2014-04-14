package main

import (
  b64 "encoding/base64"
  "fmt"
)

func main() {

  p := fmt.Println
  data := "123abc!?$*&()'-=@~"

  sEnc := b64.StdEncoding.EncodeToString([]byte(data))
  p(sEnc)

  dEnc, _ := b64.StdEncoding.DecodeString(sEnc)
  p(string(dEnc))

}
