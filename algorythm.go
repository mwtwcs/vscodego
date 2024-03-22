package main

import (
  "fmt"
)

func algorythm() {
  var a int
  fmt.Scan(&a)

  switch {
  case a%2 != 0:
    fmt.Println("YES") // Odd numbers
  case a >= 2 && a <= 5:
    fmt.Println("NO") // Range 2 to 5
  case a >= 6 && a <= 20:
    fmt.Println("YES") // Range 6 to 20
  default:
    if a%2 == 0 {
      fmt.Println("NO") // Even numbers greater than 20
    } else {
      fmt.Println("Invalid Input") // Odd numbers outside all ranges
    }
  }
}