package main

import (
  "fmt"
  "os"
)

funt main() {

  f := createFile("/tmp/defer.txt")
  defer closeFile(f)
  writeFile(f)

}
