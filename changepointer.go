package main

import "fmt"

func change(x int) {
    x = 22
}

func changePointer(pointer *int) {
    *pointer = 22
}

func pointer() {
    num := 14

    change(num)
    fmt.Println(num) // выведет 14

    changePointer(&num)
    fmt.Println(num) // выведет 22
}
