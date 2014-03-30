package main

import (
	"os"
)

func main() {

	panic("A problem has arisen")

	_, err := os.Create("/tmp/file")

	if err != nil {
		panic(err)
	}
}
