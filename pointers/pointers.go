package main

import "fmt"

func zeroval(ival int){
  ival = 0
}

func zeroptr(iptr *int){
  *iptr = 0
}

func main() {

  i := 1
  fmt.Println("initial i:", i)

  // zeroval is passed copy of i
  zeroval(i)
  fmt.Println("zeroval i:", i)

  // zeroptr dereferences i in memory, changing the value in memory
  zeroptr(&i)
  fmt.Println("zeroptr i:", i)

  fmt.Println("pointer:", &i)
}
