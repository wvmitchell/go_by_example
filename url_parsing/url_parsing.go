package main

import (
  "fmt"
  "net/url"
  "strings"
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
  p("user:", u.User)
  p("username:", u.User.Username())

  pw, _ := u.User.Password()
  p("password:", pw)

  p("host:", u.Host)

  // split on ':' to extract port
  h := strings.Split(u.Host, ":")
  p(h[0])
  p(h[1])

  // extract path and fragment
  p("path:", u.Path)
  p("fragment:", u.Fragment)

  // extract query
  p("raw query:", u.RawQuery)

  // parse query into a map
  m, _ := url.ParseQuery(u.RawQuery)
  p("query map:", m)
  p("single value:", m["k"][0])

}
