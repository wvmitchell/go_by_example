package main

import (
  "fmt"
  "time"
)

func main() {

  now := time.Now()

  // seconds since unix epoch
  secs := now.Unix()
  fmt.Println(secs)

  // nanoseconds since unix epoch
  nsecs := now.UnixNano()
  fmt.Println(nsecs)

  // milliseconds since unix epoch
  // there's no built-in function, so this must be done manually
  msecs := nsecs / 1000000
  fmt.Println(msecs)

  fmt.Println(time.Unix(secs, 0))
  fmt.Println(time.Unix(0, nsecs))

}
