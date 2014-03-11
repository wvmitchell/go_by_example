package main

import "fmt"

type person struct {
  name string
  age int
}

func main() {

  fmt.Println(person{"Bob", 20})

  fmt.Println(person{name: "Alice", age: 22})

  fmt.Println(person{age: 33, name: "Jane"})

  fmt.Println(&person{"Ann", 33})

  s := person{"Sam", 21}

  fmt.Println(s.name)

  sp := &s
  sn := s

  fmt.Println(sp.age)
  fmt.Println(sn.age)

  sp.age = 55
  fmt.Println(sp.age)
  fmt.Println(s.age)
  fmt.Println(sn.age)

}
