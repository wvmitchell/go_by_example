package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r *rect) perimeter() int {
	return 2*r.width + 2*r.height
}

func main() {

	r := rect{10, 13}
	fmt.Println("rectangle area:", r.area())
	fmt.Println("rectangle perimeter:", r.perimeter())

  rp := &r
  rp.width += 30
	fmt.Println("rectangle area:", rp.area())
	fmt.Println("rectangle perimeter:", rp.perimeter())

	fmt.Println("rectangle area:", r.area())
	fmt.Println("rectangle perimeter:", r.perimeter())

}
