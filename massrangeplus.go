
package main

import (
	"fmt"
)

func massrangeplus(){
	sum := 0
	nums := [5]int{1, 2, 3, 4, 5}
	
    for _,v := range nums{
		sum +=v
	}
	fmt.Print(sum) // Правильный ответ будет: 15
}