package main

import "fmt"

func ifelse(){
    var x1, x2 int
    fmt.Scan(&x1, &x2)

    if x1 <= 2 * x2 {
        fmt.Print("ДА")
    } else {
        fmt.Print("НЕТ")
    }
}