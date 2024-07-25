package main

import (
	"fmt"
	"os"
)

func echo1() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(os.Args[0])
}
