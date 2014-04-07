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

  // test against other times
  fp(then.Before(now))
  fp(then.After(now))
  fp(then.Equal(now))

  // Sub returns a Duration between two times
  dur := now.Sub(then)
  fp(dur)

  // Units of duration
  fp(dur.Hours())
  fp(dur.Minutes())
  fp(dur.Seconds())
  fp(dur.Nanoseconds())

  // Movining time by a given duration
  fp(then.Add(dur))
  fp(then.Add(-dur))
}
