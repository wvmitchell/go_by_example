package main

import "fmt"
import "strings"

func factorial(i int) int {
  if i == 1 {
    return 1
  } else {
    return i * factorial(i-1)
  }
}

func palindrome(s string) bool {
  if len(s) == 1 || len(s) == 0 {
    return true
  } else {
    chars := strings.Split(s, "")
    return chars[0] == chars[len(chars)-1] && palindrome(strings.Join(chars[1:len(chars)-1], ""))
  }
}

func main() {
  fmt.Println(factorial(5))
  fmt.Println(palindrome("racecar"))
  fmt.Println(palindrome("battleship"))
}
