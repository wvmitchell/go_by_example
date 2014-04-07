package main

import (
  "fmt"
  "time"
)

func main() {

  fp := fmt.Println

  // current time
  now := time.Now()
  fp(now)

  // build a time
  then := time.Date(2009, 11, 17, 20, 34, 58, 8782832, time.UTC)
  fp(then)

  // extracting information
  fp(then.Year())
  fp(then.Month())
  fp(then.Day())
  fp(then.Weekday())
  fp(then.Hour())
  fp(then.Minute())
  fp(then.Second())
  fp(then.Nanosecond())
  fp(then.Location())
}
