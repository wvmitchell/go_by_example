package main

import "fmt"

func vals() (int, int){
  return 2, 6
}

func three_squares(a int, b int, c int) (int, int, int) {
  one := a*a
  two := b*b
  three := c*c

  return one, two, three
}

func main() {
  a, b, c := three_squares(1,5,8)

  fmt.Println(a,b,c)
}
