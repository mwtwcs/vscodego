//мое решение

package main

import (
	"fmt"
)
func bit() {
	var bit float64
	fmt.Scan(&bit)
	kb := float64(bit) / 8 / 1024
	fmt.Println(kb)
}



//решение учителя
//  package bit(){}

// import ("fmt"; "math")

//func main() {
 //   var b float64
 //   fmt.Scan(&b) // 8192 bits
    
 //   fmt.Print(b / math.Pow(2, 13)) // Результат: 8192 / 2^13   
//}
