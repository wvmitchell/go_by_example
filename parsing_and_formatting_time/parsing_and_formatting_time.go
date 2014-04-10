package main

import (
  "time"
  "fmt"
)

func main() {
  p := fmt.Println

  t := time.Now()

  p(t)

  // reference based formatting
  p(t.Format("2006-01-02T15:04:05Z07:00"))
  p(t.Format("3:04PM"))
  p(t.Format("Mon Jan _2 15:04:05 2006"))
  p(t.Format("January 2 2006"))

  // standard string formatting with extracted info
  fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
    t.Year(), t.Month(), t.Day(),
    t.Hour(), t.Minute(), t.Second(),
  )
}
