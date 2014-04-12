package main

import (
  "fmt"
  "net/url"
  //"strings"
)

func main() {

  p := fmt.Println

  s := "postgres://user:pass@host.com:5432/path?k=v#f"

  u, err := url.Parse(s)

  if err != nil {
    panic(err)
  }

  p("url:", u)
  p("scheme:", u.Scheme)

}
