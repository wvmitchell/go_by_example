package main

import "fmt"
import "math/rand"

func main() {

  s := make([]string, 3)
  s[0] = "a"
  s[1] = "b"
  s[2] = "c"
  s = append(s, "boat")
  s = append(s,"d", "e")

  //fmt.Println(s)

  c := make([]string, len(s))
  copy(c,s)
  //fmt.Println(c)

  sl1 := c[2:4]
  sl2 := c[1:5]
  sl3 := c[2:]
  sl4 := c[:5]
  fmt.Println("sl1: ", sl1)
  fmt.Println("sl2: ", sl2)
  fmt.Println("sl3: ", sl3)
  fmt.Println("sl4: ", sl4)

  t := []string{"one", "two", "three"}
  fmt.Println(t)
  fmt.Println(t[:2])

  twoD := make([][]int, 4)

  for i := 0; i < len(twoD); i++ {
    inner_len := i + 1
    twoD[i] = make([]int, inner_len)
    for j := 0; j < inner_len; j++ {
      twoD[i][j] = rand.Int()%13
    }
  }

  fmt.Println(twoD)

}
